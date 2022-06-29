package dns

import (
	"encoding/base64"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	"os"
	"strings"
	"testing"
)

func TestClient(t *testing.T) {
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	client := NewClient(logger, DefaultConfig())
	userId := "lifei"
	fileId := "myfileid"
	podURI := "http://lifei.pod.fuxi./hehe/index.html"
	privkey := secp256k1.GenPrivKey()
	pubKeyBase64 := base64.StdEncoding.EncodeToString(privkey.PubKey().Bytes())

	//user register
	t.Log("user register==============================")
	err := client.RegisterUser(userId, pubKeyBase64)
	require.NoError(t, err)

	//pod register
	t.Log("pod register==============================")
	err = client.RegisterPOD(userId, podURI)
	require.NoError(t, err)

	//data upload
	t.Log("data register==============================")
	err = client.Upload(fileId, userId, podURI, "hash hex", "sha256")
	require.NoError(t, err)
}

func TestHEHE(t *testing.T) {
	userId := "lifei.user.fuxi."
	if strings.Contains(userId, "user.fuxi.") {
		userId = strings.ReplaceAll(userId, "user.fuxi.", "pod.fuxi.")
	}
	t.Log(userId)
}
