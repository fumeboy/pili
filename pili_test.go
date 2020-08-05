package pili

import (
	"testing"
)

func Test1(t *testing.T) {
	var (
		function_name = SERIAL(
			NOTBLANK(),
			TOKEN("function_name",nil),
			OPTION(SERIAL(DELIMITER(" "))),
		)
		param = PARALLEL(
			SERIAL(
				DELIMITER("'"),
				TOKEN("param", nil),
				DELIMITER("'"),
			),
			SERIAL(
				NOTBLANK(),
				TOKEN("param", nil),
				OPTION(SERIAL(DELIMITER(" "))),
			),
		)

		sentence = SERIAL(
			DELIMITER("let"),
			function_name,
			OPTION(
				SERIAL(
					REPEAT(
						"(",
						SERIAL(param),
						",",
						")",
					),
				),
			),
		)

		text = "let apple"
		ctx = NEWstate(nil,text)

		text2 = ("let apple(a,b,c,'abc','a b c')")
		ctx2 = NEWstate(nil,text2)
		//
		//text3 = ("let apple (a,b ,c,'abc','a b c')")
		//ctx3 = NEWcontext(nil,text3)
		//
		//text4 = ("let apple banana(a,b ,c,'abc','a b c')")
		//ctx4 = NEWcontext(nil,text4)
	)
	RUN(ctx,sentence)
	RUN(ctx2,sentence)
}

