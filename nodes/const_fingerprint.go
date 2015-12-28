// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node Const) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("Const")

	if node.Constbyval {
		ctx.WriteString("constbyval")
		ctx.WriteString(strconv.FormatBool(node.Constbyval))
	}

	if node.Constcollid != 0 {
		ctx.WriteString("constcollid")
		ctx.WriteString(strconv.Itoa(int(node.Constcollid)))
	}

	if node.Constisnull {
		ctx.WriteString("constisnull")
		ctx.WriteString(strconv.FormatBool(node.Constisnull))
	}

	if node.Constlen != 0 {
		ctx.WriteString("constlen")
		ctx.WriteString(strconv.Itoa(int(node.Constlen)))
	}

	if node.Consttype != 0 {
		ctx.WriteString("consttype")
		ctx.WriteString(strconv.Itoa(int(node.Consttype)))
	}

	if node.Consttypmod != 0 {
		ctx.WriteString("consttypmod")
		ctx.WriteString(strconv.Itoa(int(node.Consttypmod)))
	}

	// Intentionally ignoring node.Location for fingerprinting

	if node.Xpr != nil {
		ctx.WriteString("xpr")
		node.Xpr.Fingerprint(ctx, "Xpr")
	}
}
