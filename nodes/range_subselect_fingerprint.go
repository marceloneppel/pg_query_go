// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node RangeSubselect) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("RangeSubselect")

	if node.Alias != nil {
		ctx.WriteString("alias")
		node.Alias.Fingerprint(ctx, "Alias")
	}

	if node.Lateral {
		ctx.WriteString("lateral")
		ctx.WriteString(strconv.FormatBool(node.Lateral))
	}

	if node.Subquery != nil {
		ctx.WriteString("subquery")
		node.Subquery.Fingerprint(ctx, "Subquery")
	}
}
