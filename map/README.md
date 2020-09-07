# Map
ここで作成するのは与えられた英文の中に登場する単語がそれぞれいくつ含まれるかをmapで返す関数です。

例えば、

I am learning Go!

という英文に対しては

{"Go!":1, "I":1, "am":1, "learning":1}

を返します。

今回も"golang.org/x/tour/wc"からwc.Test()の関数を与えられています。これは作成する関数WordCountを引数にとります。引数にとった関数にたいして、任意の英文を与え、用意していた答えと同じ返り値があればPASS(合格)します。与える英文は４つあります。４つPASSすれば合格です。

[Exercise: Maps](https://tour.golang.org/moretypes/23)