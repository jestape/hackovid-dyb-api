package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/jestape/hackovid-dyb-api/src/app/model"

	"os"
	"fmt"

	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	token "github.com/jestape/hackovid-dyb-api/src/app/contracts"

)

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	user := model.User{}
	params := mux.Vars(r)

	if err := db.Where(&model.User{PublicKey: params["pk"]}).Find(&user).Error; err != nil {
		respondError(w, http.StatusOK, "User not found") 
	} else {
		respondJSON(w, http.StatusOK, user) 
	}
	
}

func GetUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	users := []model.User{}

	db.Find(&users)
	respondJSON(w, http.StatusOK, users)
	
}

func CreateSeller(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	
	var user model.User = model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if (!checkNotNil(user.PublicKey, "public_key", w) || !checkNotNil(user.Name, "user_name", w) ||
			!checkNotNil(user.Email, "email", w) || !checkNotNil(user.NIF, "nif", w)) { 
		return 
	}

	user.Role = "Seller"

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	Verify(user.PublicKey, "seller", w, r);

	respondJSON(w, http.StatusCreated, user)

}

func CreateBuyer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	var user model.User = model.User{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	
	if (!checkNotNil(user.PublicKey, "public_key", w) || !checkNotNil(user.Name, "user_name", w) ||
			!checkNotNil(user.Email, "email", w) || !checkNotNil(user.NIF, "nif", w)) { 
		return 
	}

	user.Role = "Buyer"

	if err := db.Save(&user).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	Verify(user.PublicKey, "Buyer", w, r);

	respondJSON(w, http.StatusCreated, user)

}

func Verify(public_key string  , types string, w http.ResponseWriter, r *http.Request) {

	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/" + os.Getenv("INFURA_PROJECT_ID"))
    if err != nil {
        respondError(w, http.StatusBadRequest, "ethClient: " + err.Error())
		return
	}
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
    if err != nil {
		respondError(w, http.StatusBadRequest, "Private Key : " + err.Error())
		return
	}
	publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
		respondError(w, http.StatusBadRequest, "cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		respondError(w, http.StatusBadRequest, "From: " + err.Error())
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
		respondError(w, http.StatusBadRequest, "GasPrice: " + err.Error())
		return
	}

	auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)
    auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	
	address := common2.HexToAddress(os.Getenv("SMART_CONTRACT_ADDRESS"))
    instance, err := token.NewStore(address, client)
    if err != nil {
		respondError(w, http.StatusBadRequest, "Smart Contract: " + err.Error())
		return
	}

	user := common2.HexToAddress(public_key)

	if (types == "seller") {
		tx, err := instance.AddSeller(auth, user);
		if err != nil {
			respondError(w, http.StatusBadRequest, "Transaction: " + err.Error())
			return
		}
		fmt.Printf("tx sent: %s", tx.Hash().Hex())
	} else {
		tx, err := instance.AddBuyer(auth, user);
		if err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		}
		fmt.Printf("tx sent: %s", tx.Hash().Hex())
	}
}

