package controllers

import (
	"encoding/json"
	"redis/models"
	"redis/utils"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// GetEmails get emails
func GetEmails() ([]models.Emails, error) {

	var emails []models.Emails
	client := utils.GetRedisClient()

	keys, err := client.Keys("*").Result()
	if err != nil {
		log.Error(err)
		return []models.Emails{}, err
	}

	for _, k := range keys {
		var email models.Emails

		resp, err := client.Get(k).Result()
		if err != nil {
			log.Error(err)
			return []models.Emails{}, err
		}

		b := []byte(resp)
		err = json.Unmarshal(b, &email)
		if err != nil {
			log.Error(err)
			return []models.Emails{}, err
		}

		emails = append(emails, email)
	}

	return emails, nil
}

func AddEmail(request models.Emails) (bool, error) {
	client := utils.GetRedisClient()

	result, err := json.Marshal(request)
	if err != nil {
		log.Error(err)
		return false, err
	}

	_, err = client.Do("SET", viper.GetString("redis.prefix")+request.Email, string(result)).Result()
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}

func GetEmail(param string) (models.Emails, error) {
	client := utils.GetRedisClient()

	var email models.Emails

	resp, err := client.Get(viper.GetString("redis.prefix") + param).Result()
	if err != nil {
		log.Error(err)
		return models.Emails{}, err
	}

	b := []byte(resp)
	err = json.Unmarshal(b, &email)
	if err != nil {
		log.Error(err)
		return models.Emails{}, err
	}

	return email, nil
}

func UpdateEmail(request models.Emails, param string) (bool, error) {
	client := utils.GetRedisClient()

	var currentEmail, email models.Emails

	resp, err := client.Get(viper.GetString("redis.prefix") + param).Result()
	if err != nil {
		log.Error(err)
		return false, err
	}

	b := []byte(resp)
	err = json.Unmarshal(b, &currentEmail)
	if err != nil {
		log.Error(err)
		return false, err
	}

	email.FirstName = request.FirstName
	email.LastName = request.LastName
	email.Email = request.Email

	result, err := json.Marshal(email)
	if err != nil {
		log.Error(err)
		return false, err
	}

	_, err = client.Rename(viper.GetString("redis.prefix")+param, viper.GetString("redis.prefix")+email.Email).Result()
	if err != nil {
		log.Error(err)
		return false, err
	}

	_, err = client.Do("SET", viper.GetString("redis.prefix")+email.Email, string(result)).Result()
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}

func DeleteEmail(param string) (bool, error) {
	client := utils.GetRedisClient()

	_, err := client.Get(viper.GetString("redis.prefix") + param).Result()
	if err != nil {
		log.Error(err)
		return false, err
	}

	_, err = client.Del(viper.GetString("redis.prefix") + param).Result()
	if err != nil {
		log.Error(err)
		return false, err
	}

	return true, nil
}
