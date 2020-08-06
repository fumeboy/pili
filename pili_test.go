package pili

import (
	"testing"
)

func Test1(t *testing.T) {
	var (
		function_name = SERIAL(
			NOTBLANK(),
			TOKEN("function_name",nil),
			OPTION(SERIAL(BLANK())),
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
				OPTION(SERIAL(BLANK())),
			),
		)

		sentence = SERIAL(
			DELIMITER("let"),
			BLANK(),
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

		text3 = ("letapple (a,b ,c,'abc' ,'a b c')")
		ctx3 = NEWstate(nil,text3)
	)
	ctx.RUN(sentence)
	ctx2.RUN(sentence)
	ctx3.RUN(sentence)
}

