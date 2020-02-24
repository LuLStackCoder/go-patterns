package main

import (
	`fmt`

	"github.com/LuLStackCoder/go-patterns/pkg/tex_doc"
	"github.com/LuLStackCoder/go-patterns/pkg/html_doc"
	"github.com/LuLStackCoder/go-patterns/pkg/visitor"
)

const (
	htmlString = `<td class="played">0</td>`
	texString1 = `The quadratic formula is $-b \pm \sqrt{b^2 - 4ac} \over 2a$
\bye`
	texString2 = `\begin{document}

\section{Hello \textit{world}.}

\subsection{Watermelon}

(n.) A sacred fruit. Also known as:

\begin{itemize}
\item red lemon
\item life
\end{itemize}

Here is the prevalence of each synonym.

\begin{tabular}{c c}
red lemon & uncommon \\
life & common
\end{tabular}

\end{document}`
)

func main() {
	var newVisitor = visitor.NewVisitor()
	var newHtml = html_doc.NewHtmlDoc(htmlString)
	var newTex1 = tex_doc.NewTexDoc(texString1)
	var newTex2 = tex_doc.NewTexDoc(texString2)
	fmt.Println(newHtml.Accept(newVisitor))
	fmt.Println(newTex1.Accept(newVisitor))
	fmt.Println(newTex2.Accept(newVisitor))
}