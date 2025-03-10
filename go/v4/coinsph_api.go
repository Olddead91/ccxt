// -------------------------------------------------------------------------------

// PLEASE DO NOT EDIT THIS FILE, IT IS GENERATED AND WILL BE OVERWRITTEN:
// https://github.com/ccxt/ccxt/blob/master/CONTRIBUTING.md#how-to-contribute-code

// -------------------------------------------------------------------------------

package ccxt

func (this *coinsph) PublicGetOpenapiV1Ping (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("publicGetOpenapiV1Ping", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PublicGetOpenapiV1Time (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("publicGetOpenapiV1Time", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PublicGetOpenapiQuoteV1Ticker24hr (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("publicGetOpenapiQuoteV1Ticker24hr", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PublicGetOpenapiQuoteV1TickerPrice (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("publicGetOpenapiQuoteV1TickerPrice", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PublicGetOpenapiQuoteV1TickerBookTicker (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("publicGetOpenapiQuoteV1TickerBookTicker", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PublicGetOpenapiV1ExchangeInfo (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("publicGetOpenapiV1ExchangeInfo", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PublicGetOpenapiQuoteV1Depth (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("publicGetOpenapiQuoteV1Depth", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PublicGetOpenapiQuoteV1Klines (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("publicGetOpenapiQuoteV1Klines", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PublicGetOpenapiQuoteV1Trades (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("publicGetOpenapiQuoteV1Trades", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PublicGetOpenapiV1Pairs (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("publicGetOpenapiV1Pairs", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PublicGetOpenapiQuoteV1AvgPrice (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("publicGetOpenapiQuoteV1AvgPrice", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetOpenapiWalletV1ConfigGetall (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetOpenapiWalletV1ConfigGetall", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetOpenapiWalletV1DepositAddress (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetOpenapiWalletV1DepositAddress", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetOpenapiWalletV1DepositHistory (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetOpenapiWalletV1DepositHistory", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetOpenapiWalletV1WithdrawHistory (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetOpenapiWalletV1WithdrawHistory", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetOpenapiV1Account (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetOpenapiV1Account", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetOpenapiV1OpenOrders (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetOpenapiV1OpenOrders", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetOpenapiV1AssetTradeFee (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetOpenapiV1AssetTradeFee", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetOpenapiV1Order (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetOpenapiV1Order", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetOpenapiV1HistoryOrders (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetOpenapiV1HistoryOrders", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetOpenapiV1MyTrades (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetOpenapiV1MyTrades", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetOpenapiV1CapitalDepositHistory (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetOpenapiV1CapitalDepositHistory", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetOpenapiV1CapitalWithdrawHistory (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetOpenapiV1CapitalWithdrawHistory", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetOpenapiV3PaymentRequestGetPaymentRequest (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetOpenapiV3PaymentRequestGetPaymentRequest", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetMerchantApiV1GetInvoices (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetMerchantApiV1GetInvoices", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetOpenapiAccountV3CryptoAccounts (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetOpenapiAccountV3CryptoAccounts", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateGetOpenapiTransferV3TransfersId (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateGetOpenapiTransferV3TransfersId", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiWalletV1WithdrawApply (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiWalletV1WithdrawApply", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiV1OrderTest (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiV1OrderTest", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiV1Order (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiV1Order", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiV1CapitalWithdrawApply (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiV1CapitalWithdrawApply", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiV1CapitalDepositApply (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiV1CapitalDepositApply", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiV3PaymentRequestPaymentRequests (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiV3PaymentRequestPaymentRequests", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiV3PaymentRequestDeletePaymentRequest (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiV3PaymentRequestDeletePaymentRequest", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiV3PaymentRequestPaymentRequestReminder (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiV3PaymentRequestPaymentRequestReminder", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiV1UserDataStream (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiV1UserDataStream", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostMerchantApiV1Invoices (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostMerchantApiV1Invoices", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostMerchantApiV1InvoicesCancel (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostMerchantApiV1InvoicesCancel", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiConvertV1GetSupportedTradingPairs (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiConvertV1GetSupportedTradingPairs", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiConvertV1GetQuote (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiConvertV1GetQuote", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiConvertV1AccpetQuote (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiConvertV1AccpetQuote", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiFiatV1SupportChannel (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiFiatV1SupportChannel", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiFiatV1CashOut (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiFiatV1CashOut", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiFiatV1History (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiFiatV1History", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiMigrationV4Sellorder (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiMigrationV4Sellorder", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiMigrationV4ValidateField (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiMigrationV4ValidateField", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivatePostOpenapiTransferV3Transfers (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privatePostOpenapiTransferV3Transfers", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateDeleteOpenapiV1Order (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateDeleteOpenapiV1Order", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateDeleteOpenapiV1OpenOrders (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateDeleteOpenapiV1OpenOrders", parameters))
       PanicOnError(ch)
   }()
   return ch
}

func (this *coinsph) PrivateDeleteOpenapiV1UserDataStream (args ...interface{}) <-chan interface{} {
   parameters := GetArg(args, 0, nil)
   ch := make(chan interface{})
   go func() {
       defer close(ch)
       defer func() {
           if r := recover(); r != nil {
               ch <- "panic:" + ToString(r)
           }
       }()
       ch <- (<-this.callEndpoint ("privateDeleteOpenapiV1UserDataStream", parameters))
       PanicOnError(ch)
   }()
   return ch
}
