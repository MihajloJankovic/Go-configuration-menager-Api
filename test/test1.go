package test

//Zakomentarisani importi su prekopirani iz Dao\Dao.go i Dao\helper.go jer se func iz ta dva testiraju.
//Obrisati one koje ne budu potrebne
import (
	"testing"
	//"context"
	//"encoding/json"
	//"fmt"
	//tracer "github.com/MihajloJankovic/Alati/tracer"
	//"github.com/hashicorp/consul/api"
	//"os"
	//"github.com/google/uuid"
)

// Test bez mockinga
func TestOne(t *testing.T) {
	//predefinisati parametre i ocekivane rezultate za neku funkciju iz Dao.go(Get bi verovatno bila najlaksa)
	//pokrenuti funkciju
	//u slucaju da se ocekivani i stvarni rezultati ne poklapaju, ispis se vrsi sa t.Errorf
	//t.Errorf("Test failed. Expected: %d, Actual: %d", expected, actual)
}

// Test sa mockingom
func TestTwo(t *testing.T) {

}
