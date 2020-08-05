看似 parser combinator 但其实完全不是的又一个原创分词器

串行 SERIAL
并行 PARALLEL
可选并行 OPTION
重复匹配 REPEAT

定界符 DELIMITER


人工部分：
    使用 SERIAL、PARALLEL、TOKEN、EVENT、DELIMITER、REPEAT 这些工具书写 期望匹配的 句型
程序自动部分：
    产生 routes 树，每个结点持有一个 matcher 函数用以匹配文本
执行部分：
    输入文本，根据 routes 树进行文本匹配，并从输入中取出目的文本