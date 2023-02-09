package main

import (
	"fmt"
	"github.com/open-wallstreet/yahoo-live/pkg/yahoo"
	"github.com/open-wallstreet/yahoo-live/pkg/yahoo/proto"
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewDevelopment()
	con, err := yahoo.NewWebsocket(logger.Sugar(), []string{
		"^GSPC",
		"^DJI",
		"^IXIC",
		"^RUT",
		"CL=F",
		"GC=F",
		"KAMBI.ST",
		"KIND-SDB.ST",
		"AKBNK.IS",
		"AKBTY",
		"DIS",
		"AFRM",
		"HOOD",
		"GOOG",
		"WYNN",
		"TSLA",
		"LUMN",
		"UBER",
		"GOOGL",
		"BTC-USD",
		"ETH-USD",
		"USDT-USD",
		"BNB-USD",
		"USDC-USD",
		"CPRI",
		"AUKUF",
		"ADYEY",
		"VSAT",
		"KD",
		"NEWR",
		"WFRD",
		"NYT",
		"SNEX",
		"MFLX",
		"NFTY",
		"FLIN",
		"IHF",
		"FLN",
		"RYSIX",
		"RYSCX",
		"BIOIX",
		"BIOPX",
		"BIOUX",
		"ESH23.CME",
		"YMH23.CBT",
		"NQH23.CME",
		"RTYH23.CME",
		"ZBH23.CBT",
		"SKLZ",
		"FUBO",
		"EXPR",
		"NCMI",
		"GOTU",
		"EURUSD=X",
		"JPY=X",
		"GBPUSD=X",
		"AUDUSD=X",
		"NZDUSD=X",
		"AMC",
		"QQQ",
		"EEM",
		"HYG",
	})
	if err != nil {
		panic(err)
	}
	con.AddMessageHandler(on_msg)
	con.Wait()
}
func on_msg(message *proto.Yaticker) {
	println(fmt.Sprintf("%s: %s", time.Unix(message.Time/1000, 0).String(), message.String()))
}
