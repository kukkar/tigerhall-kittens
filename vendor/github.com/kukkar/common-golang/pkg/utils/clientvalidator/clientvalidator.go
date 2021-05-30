package clientvalidator

import (
	"context"
	"fmt"

	"github.com/kukkar/common-golang/pkg/utils/aesencryption"
	"github.com/kukkar/common-golang/pkg/utils/hashmac"
	"github.com/kukkar/common-golang/pkg/utils/hmacrequestgenerator"
	"github.com/kukkar/common-golang/pkg/utils/rError"
)

//
// ValidateHmac function take request hmac and client as param and verify the client
// after parsing the request from aes and hmac algo
//
func ValidateHmac(requestPayload map[string]interface{}, hmac string,
	clientName string, clientValidator ClientValidator) (bool, error) {

	c := context.TODO()
	generatedHMac, err := GenerateHMac(c, requestPayload,
		clientName, clientValidator)
	if err != nil {
		err = rError.MiscError(c, err, "unable to generate Hmac")
		return false, err
	}
	base64HMac, err := GenerateBase64HMac(c, requestPayload,
		clientName, clientValidator)
	if err != nil {
		err = rError.MiscError(c, err, "unable to generate Hmac")
		return false, err
	}
	if fmt.Sprintf("%x", generatedHMac) != hmac {
		if fmt.Sprintf("%s", base64HMac) != hmac {
			err := rError.UnauthoriseErr(c, fmt.Sprintf("unauthorised Client %s with hmac %s", clientName, hmac))
			return false, err
		}
	}

	return true, nil
}

//
// GenerateHMac
//
func GenerateHMac(c context.Context, requestPayload map[string]interface{},
	clientName string, clientValidator ClientValidator) ([]byte, error) {

	// //merchant sdk to reterive merchant master key and master keys of system
	// merchantSDK, err := merchant.GetMerchantSdk(merchantDBName)
	// if err != nil {
	// 	err = rError.MiscError(c, err, "not able to get merchant sdk")
	// 	return nil, err
	// }

	decryptedMasterKey, err := GetMasterSecret(c, clientValidator)
	if err != nil {
		err = rError.MiscError(c, err, "Unable to decrypt master key using super key")
		return nil, err
	}
	//aes ecnryption instance with loaded new master key
	aes4Merchant, err := aesencryption.GetAesEncyption((*decryptedMasterKey))
	if err != nil {
		err = rError.MiscError(c, err, "Unable to Load decrypted master key")
		return nil, err
	}
	// get merchant secret key
	clientSecret, err := clientValidator.GetClientSecret(clientName)
	if err != nil {
		err = rError.MiscError(c, err, fmt.Sprintf("Unable to get merchant secret key using client Name %s ", clientName))
		return nil, err
	}
	//decrypt merchant secret key using master key
	decryptedClientSecret, err := aes4Merchant.Decrypt(clientSecret)
	if err != nil {
		err = rError.MiscError(c, err, "Unable to decrypt client secret using master key")
		return nil, err
	}
	parsedHMacRequest, err := hmacrequestgenerator.GetHMacRequest(requestPayload)
	if err != nil {
		err = rError.MiscError(c, err, "unable to generate hmac request payload")
		return nil, err
	}
	generatedHMac := hashmac.GenerateMac(parsedHMacRequest, decryptedClientSecret)
	fmt.Println(fmt.Sprintf("Calculated Hmac   %x ,payload sort %s", generatedHMac, parsedHMacRequest))
	return generatedHMac, nil
}

//
// GenerateBase64HMac generate base64 hmac
//
func GenerateBase64HMac(c context.Context, requestPayload map[string]interface{},
	clientName string, clientValidator ClientValidator) ([]byte, error) {

	// //merchant sdk to reterive merchant master key and master keys of system
	// merchantSDK, err := merchant.GetMerchantSdk(merchantDBName)
	// if err != nil {
	// 	err = rError.MiscError(c, err, "not able to get merchant sdk")
	// 	return nil, err
	// }

	decryptedMasterKey, err := GetMasterSecret(c, clientValidator)
	if err != nil {
		err = rError.MiscError(c, err, "Unable to decrypt master key using super key")
		return nil, err
	}
	//aes ecnryption instance with loaded new master key
	aes4Merchant, err := aesencryption.GetAesEncyption((*decryptedMasterKey))
	if err != nil {
		err = rError.MiscError(c, err, "Unable to Load decrypted master key")
		return nil, err
	}
	// get merchant secret key
	clientSecret, err := clientValidator.GetClientSecret(clientName)
	if err != nil {
		err = rError.MiscError(c, err, fmt.Sprintf("Unable to get merchant secret key using client Name %s ", clientName))
		return nil, err
	}
	//decrypt merchant secret key using master key
	decryptedClientSecret, err := aes4Merchant.Decrypt(clientSecret)
	if err != nil {
		err = rError.MiscError(c, err, "Unable to decrypt client secret using master key")
		return nil, err
	}
	parsedHMacRequest, err := hmacrequestgenerator.GetHMacRequest(requestPayload)
	if err != nil {
		err = rError.MiscError(c, err, "unable to generate hmac request payload")
		return nil, err
	}
	generatedHMac := hashmac.GenerateBase64Mac(parsedHMacRequest, decryptedClientSecret)
	fmt.Println(fmt.Sprintf("Calculated Hmac   %s ,payload sort %s", generatedHMac, parsedHMacRequest))
	return generatedHMac, nil
}

//
// GetMasterSecret return decrypted master secret key
//
func GetMasterSecret(c context.Context, clientValidator ClientValidator) (*string, error) {

	if MasterKeyDecrypted != nil {
		return MasterKeyDecrypted, nil
	}
	// //merchant sdk to reterive merchant master key and master keys of system
	// merchantSDK, err := merchant.GetMerchantSdk(merchantDBName)
	// if err != nil {
	// 	err = rError.MiscError(c, err, "not able to get merchant sdk")
	// 	return nil, err
	// }

	superKey, err := clientValidator.GetSuperKey()
	if err != nil {
		return nil, err
	}
	//aes ecnryption instance with loaded super key
	aesEncrypted, err := aesencryption.GetAesEncyption(superKey)
	if err != nil {
		err = rError.MiscError(c, err, "failed to get aes encryption instance with super key")
		return nil, err
	}
	masterSecret, err := clientValidator.GetMasterKey()
	if err != nil {
		err = rError.MiscError(c, err, "unable to fetch master key")
		return nil, err
	}
	decryptedMasterKey, err := aesEncrypted.Decrypt(masterSecret)
	if err != nil {
		err = rError.MiscError(c, err, "Unable to decrypt master key using super key")
		return nil, err
	}
	MasterKeyDecrypted = &decryptedMasterKey
	return MasterKeyDecrypted, nil
}
