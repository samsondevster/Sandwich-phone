func generateRandomUsername(r *rand.Rand) (ret string) {
	if r.Int63n(100) > 9 {
		ret, _ = regen.Generate("^[A-Za-z0-9]+(?:[ _-][A-Za-z0-9]+)*$")
	} else {
		ret = invalidField
	}
	return
}

func generateRandomEmail(r *rand.Rand) (ret string) {
	if r.Int63n(100) > 9 {
		ret, _ = regen.Generate("^([a-zA-Z0-9_\\-.]+)@([a-zA-Z0-9_\\-.]+)\\.([a-zA-Z0-9]{2,})$")
	} else {
		ret = invalidField
	}
	return
}

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	/*
		return []simtypes.ParamChange{
			simulation.NewSimParamChange(types.ModuleName, string(types.ParamStoreKeyItemTransferFeePercentage),
				func(r *rand.Rand) string {

					return "TODO"
				},
			),

func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
  		case bytes.Equal(kvA.Key[:1], types.KeyPrefix(types.GoogleInAppPurchaseOrderKey)):
			var giapA, giapB types.GoogleInAppPurchaseOrder
			cdc.MustUnmarshal(kvA.Value, &giapA)
			cdc.MustUnmarshal(kvB.Value, &giapB)
			return fmt.Sprintf("%v\n%v", giapA, giapB)}
