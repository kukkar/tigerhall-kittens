package otp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/kukkar/common-golang/globalconst"
	"github.com/kukkar/common-golang/pkg/logger"
	"github.com/kukkar/common-golang/pkg/utils/clientvalidator"
)

//OTPAccessor
type OTPAccessor struct {
	IPPort          string
	Version         string
	SuperKey        string
	MerchantDBToUse string
	clientvalidator clientvalidator.ClientValidator
}

func (this OTPAccessor) VerfiyOTP(otp string, uuid string,
	clientName string) error {

	var res verifyOTPRes

	re := verifyOTPReq{
		UUID: uuid,
		OTP:  otp,
	}
	url := fmt.Sprintf("%s%s/%s", this.IPPort, this.Version, VerfiyOTPRoute)
	j, err := json.Marshal(re)
	if err != nil {
		return err
	}
	var hmacReq map[string]interface{}
	err = json.Unmarshal(j, &hmacReq)
	if err != nil {
		return err
	}
	c := context.TODO()
	hmac, err := clientvalidator.GenerateHMac(c, hmacReq,
		clientName, this.clientvalidator)
	if err != nil {
		return fmt.Errorf(fmt.Sprintf("Generate hmac failed with Error %v", err))
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	req.Header.Add(globalconst.HMAC, fmt.Sprintf("%x", hmac))
	req.Header.Add(globalconst.CLIENT_NAME, clientName)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if 200 != resp.StatusCode {
		return fmt.Errorf("%s", body)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}
	return nil
}

func (this OTPAccessor) GenerateOTP(otpReq GenerateOTPReq) (*GenerateOTPResponse, error) {

	var res generateOTPServiceResponse
	var output GenerateOTPResponse

	url := fmt.Sprintf("%s%s/%s", this.IPPort, this.Version, GenerateOTPRoute)
	j, err := json.Marshal(otpReq)
	if err != nil {
		return nil, err
	}
	var hmacReq map[string]interface{}
	err = json.Unmarshal(j, &hmacReq)
	if err != nil {
		return nil, err
	}
	c := context.TODO()
	hmac, err := clientvalidator.GenerateHMac(c, hmacReq,
		otpReq.ClientName, this.clientvalidator)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Add(globalconst.HMAC, fmt.Sprintf("%x", hmac))
	req.Header.Add(globalconst.CLIENT_NAME, otpReq.ClientName)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	copier.Copy(&output, &res.Data)
	return &output, nil
}

func (this OTPAccessor) ResendOTP(otpReq ResendOTPReq) error {

	var res generateOTPServiceResponse
	var output ResendOTPServiceRes

	url := fmt.Sprintf("%s%s/%s", this.IPPort, this.Version, ResendOTPRoute)
	j, err := json.Marshal(otpReq)
	if err != nil {
		return err
	}
	var hmacReq map[string]interface{}
	err = json.Unmarshal(j, &hmacReq)
	if err != nil {
		return err
	}
	c := context.TODO()
	hmac, err := clientvalidator.GenerateHMac(c, hmacReq,
		otpReq.ClientName, this.clientvalidator)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	req.Header.Add(globalconst.HMAC, fmt.Sprintf("%x", hmac))
	req.Header.Add(globalconst.CLIENT_NAME, otpReq.ClientName)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if 200 != resp.StatusCode {
		return fmt.Errorf("%s", body)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err
	}
	copier.Copy(&output, &res.Data)
	logger.Logger.Info(fmt.Sprintf("output %v", output))
	return nil
}

func (this OTPAccessor) VerfiyOTPV2(otp string, uuid string,
	clientName string) (*VerifyOtpData, error) {

	var res VerifyOtpResV2

	re := verifyOTPReq{
		UUID: uuid,
		OTP:  otp,
	}
	url := fmt.Sprintf("%s%s/%s", this.IPPort, "v2", VerfiyOTPRoute)
	j, err := json.Marshal(re)
	if err != nil {
		return nil, err
	}
	var hmacReq map[string]interface{}
	err = json.Unmarshal(j, &hmacReq)
	if err != nil {
		return nil, err
	}
	c := context.TODO()
	hmac, err := clientvalidator.GenerateHMac(c, hmacReq,
		clientName, this.clientvalidator)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Generate hmac failed with Error %v", err))
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Add(globalconst.HMAC, fmt.Sprintf("%x", hmac))
	req.Header.Add(globalconst.CLIENT_NAME, clientName)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return &res.Data, nil
}
