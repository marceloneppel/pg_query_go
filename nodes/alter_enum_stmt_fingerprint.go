// Auto-generated - DO NOT EDIT

package pg_query

import "strconv"

func (node AlterEnumStmt) Fingerprint(ctx FingerprintContext, parentFieldName string) {
	ctx.WriteString("AlterEnumStmt")

	if node.NewVal != nil {
		ctx.WriteString("newVal")
		ctx.WriteString(*node.NewVal)
	}

	if node.NewValIsAfter {
		ctx.WriteString("newValIsAfter")
		ctx.WriteString(strconv.FormatBool(node.NewValIsAfter))
	}

	if node.NewValNeighbor != nil {
		ctx.WriteString("newValNeighbor")
		ctx.WriteString(*node.NewValNeighbor)
	}

	if node.SkipIfExists {
		ctx.WriteString("skipIfExists")
		ctx.WriteString(strconv.FormatBool(node.SkipIfExists))
	}

	if len(node.TypeName.Items) > 0 {
		ctx.WriteString("typeName")
		node.TypeName.Fingerprint(ctx, "TypeName")
	}
}
