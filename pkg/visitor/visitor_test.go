package visitor

import (
	`testing`

	`github.com/stretchr/testify/assert`

	`github.com/LuLStackCoder/go-patterns/pkg/html_doc`
	`github.com/LuLStackCoder/go-patterns/pkg/tex_doc`
)

const (
	longHtmlElement  = `<div class="ac ae af ag ah cq aj ak"><div><div id="f26b" class="cr cs ct bk cu b cv cw cx cy cz da db dc dd de df"><h1 class="cu b cv dg cx dh cz di db dj dd dk ct"><strong class="az">Visitor Design Pattern in golang</strong></h1></div><div class="dl"><div class="n dm dn do dp"><div class="o n"><div><a rel="noopener" href="/@felipedutratine?source=post_page-----3c142a12945a----------------------"><img alt="Felipe Dutra Tine e Silva" class="r dq dr ds" src="https://miro.medium.com/fit/c/96/96/1*VbH-8ih6EfW3m5uE0x_jHg.jpeg" width="48" height="48"></a></div><div class="dt ak r"><div class="n"><div style="flex:1"><span class="bj b bk bl bm bn r ct q"><div class="du n o dv"><span class="bj dw dx bl dy dz ea eb ec ed ct"><a class="at au av aw ax ay az ba bb bc ee bf bg bh bi" rel="noopener" href="/@felipedutratine?source=post_page-----3c142a12945a----------------------">Felipe Dutra Tine e Silva</a></span><div class="ef r ar h"><button class="eg ct q eh ei ej ek el bc bh em en eo ep eq er es bj b bk et eu bn ev ew ce ex ey bf">Follow</button></div></div></span></div></div><span class="bj b bk bl bm bn r bo bp"><span class="bj dw dx bl dy dz ea eb ec ed bo"><div><a class="at au av aw ax ay az ba bb bc ee bf bg bh bi" rel="noopener" href="/@felipedutratine/visitor-design-pattern-in-golang-3c142a12945a?source=post_page-----3c142a12945a----------------------">Dec 19, 2019</a> <!-- -->·<!-- --> <!-- -->2<!-- --> min read</div></span></span></div></div><div class="n ez fa fb fc fd fe ff fg ab"><div class="n o"><div class="fh r ar"><a href="//medium.com/p/3c142a12945a/share/twitter?source=post_actions_header---------------------------" class="at au av aw ax ay az ba bb bc bd be bf bg bh bi" target="_blank" rel="noopener nofollow"><svg width="29" height="29" class="q"><path d="M22.05 7.54a4.47 4.47 0 0 0-3.3-1.46 4.53 4.53 0 0 0-4.53 4.53c0 .35.04.7.08 1.05A12.9 12.9 0 0 1 5 6.89a5.1 5.1 0 0 0-.65 2.26c.03 1.6.83 2.99 2.02 3.79a4.3 4.3 0 0 1-2.02-.57v.08a4.55 4.55 0 0 0 3.63 4.44c-.4.08-.8.13-1.21.16l-.81-.08a4.54 4.54 0 0 0 4.2 3.15 9.56 9.56 0 0 1-5.66 1.94l-1.05-.08c2 1.27 4.38 2.02 6.94 2.02 8.3 0 12.86-6.9 12.84-12.85.02-.24 0-.43 0-.65a8.68 8.68 0 0 0 2.26-2.34c-.82.38-1.7.62-2.6.72a4.37 4.37 0 0 0 1.95-2.51c-.84.53-1.81.9-2.83 1.13z"></path></svg></a></div><div class="fh r ar"><button class="at au av aw ax ay az ba bb bc bd be bf bg bh bi"><svg width="29" height="29" viewBox="0 0 29 29" fill="none" class="q"><path d="M5 6.36C5 5.61 5.63 5 6.4 5h16.2c.77 0 1.4.61 1.4 1.36v16.28c0 .75-.63 1.36-1.4 1.36H6.4c-.77 0-1.4-.6-1.4-1.36V6.36z"></path><path fill-rule="evenodd" clip-rule="evenodd" d="M10.76 20.9v-8.57H7.89v8.58h2.87zm-1.44-9.75c1 0 1.63-.65 1.63-1.48-.02-.84-.62-1.48-1.6-1.48-.99 0-1.63.64-1.63 1.48 0 .83.62 1.48 1.59 1.48h.01zM12.35 20.9h2.87v-4.79c0-.25.02-.5.1-.7.2-.5.67-1.04 1.46-1.04 1.04 0 1.46.8 1.46 1.95v4.59h2.87v-4.92c0-2.64-1.42-3.87-3.3-3.87-1.55 0-2.23.86-2.61 1.45h.02v-1.24h-2.87c.04.8 0 8.58 0 8.58z" fill="#fff"></path></svg></button></div><div class="fh r ar"><a href="//medium.com/p/3c142a12945a/share/facebook?source=post_actions_header---------------------------" class="at au av aw ax ay az ba bb bc bd be bf bg bh bi" target="_blank" rel="noopener nofollow"><svg width="29" height="29" class="q"><path d="M23.2 5H5.8a.8.8 0 0 0-.8.8V23.2c0 .44.35.8.8.8h9.3v-7.13h-2.38V13.9h2.38v-2.38c0-2.45 1.55-3.66 3.74-3.66 1.05 0 1.95.08 2.2.11v2.57h-1.5c-1.2 0-1.48.57-1.48 1.4v1.96h2.97l-.6 2.97h-2.37l.05 7.12h5.1a.8.8 0 0 0 .79-.8V5.8a.8.8 0 0 0-.8-.79"></path></svg></a></div><div class="fi r ao"><a href="https://medium.com/m/signin?operation=register&amp;redirect=https%3A%2F%2Fmedium.com%2F%40felipedutratine%2Fvisitor-design-pattern-in-golang-3c142a12945a&amp;source=post_actions_header--------------------------bookmark_sidebar-" class="at au av aw ax ay az ba bb bc bd be bf bg bh bi" rel="noopener"><svg width="25" height="25" viewBox="0 0 25 25"><path d="M19 6a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v14.66h.01c.01.1.05.2.12.28a.5.5 0 0 0 .7.03l5.67-4.12 5.66 4.13a.5.5 0 0 0 .71-.03.5.5 0 0 0 .12-.29H19V6zm-6.84 9.97L7 19.64V6a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1v13.64l-5.16-3.67a.49.49 0 0 0-.68 0z" fill-rule="evenodd"></path></svg></a></div></div></div></div></div></div><p id="31a8" class="fj fk ct bk fl b fm fn fo fp fq fr fs ft fu fv fw cl" data-selectable-paragraph="">Wikipedia says :</p><blockquote class="fx fy fz"><p id="7d52" class="fj fk ct ga fl b fm fn fo fp fq fr fs ft fu fv fw cl" data-selectable-paragraph="">In object-oriented programming and software engineering, the visitor design pattern is a way of separating an algorithm from an object structure on which it operates. A practical result of this separation is the ability to add new operations to existing object structures without modifying the structures. It is one way to follow the open/closed principle.</p></blockquote><p id="43a1" class="fj fk ct bk fl b fm fn fo fp fq fr fs ft fu fv fw cl" data-selectable-paragraph="">An example being better than 100 words, lets take an example.</p><p id="37f8" class="fj fk ct bk fl b fm fn fo fp fq fr fs ft fu fv fw cl" data-selectable-paragraph="">So we have our Employee Developer and Directors that is a structure that works well and are critical.</p><pre class="gb gc gd ge gf gg gh gi"><span id="b4ce" class="gj gk ct bk gl b dx gm gn r go" data-selectable-paragraph="">type Employee interface {<br>  FullName()<br>}</span><span id="ce91" class="gj gk ct bk gl b dx gp gq gr gs gt gn r go" data-selectable-paragraph="">type Developer struct {<br>  FirstName string<br>  LastName string<br>  Income float32<br>  Age int<br>}</span><span id="e618" class="gj gk ct bk gl b dx gp gq gr gs gt gn r go" data-selectable-paragraph="">func (d Developer) FullName() {<br>  fmt.Println(“Developer “,d.FirstName,” “,d.LastName)<br>}</span><span id="cbb2" class="gj gk ct bk gl b dx gp gq gr gs gt gn r go" data-selectable-paragraph="">type Director struct {<br>  FirstName string<br>  LastName string<br>  Income float32<br>  Age int<br>}</span><span id="bc9c" class="gj gk ct bk gl b dx gp gq gr gs gt gn r go" data-selectable-paragraph="">func (d Director) FullName() {<br>  fmt.Println(“Director “,d.FirstName,” “,d.LastName)<br>}</span></pre><p id="dc0d" class="fj fk ct bk fl b fm fn fo fp fq fr fs ft fu fv fw cl" data-selectable-paragraph="">An for some raison we want to <em class="ga">*extend*</em> that structure. But as these objects are highly critical we don’t want to change it (not too much).</p><p id="07c9" class="fj fk ct bk fl b fm fn fo fp fq fr fs ft fu fv fw cl" data-selectable-paragraph="">So We will use for that the Visitor Pattern, in order to “inject” new operations to existing object structures without modifying the structures.</p><pre class="gb gc gd ge gf gg gh gi"><span id="0fa0" class="gj gk ct bk gl b dx gm gn r go" data-selectable-paragraph="">type Visitor interface {<br>  VisitDeveloper(d Developer)<br>  VisitDirector(d Director)<br>}</span></pre><p id="f0a1" class="fj fk ct bk fl b fm fn fo fp fq fr fs ft fu fv fw cl" data-selectable-paragraph="">Now each new “additional Operation” or Behavior can be coded in a separate structure.</p><pre class="gb gc gd ge gf gg gh gi"><span id="d0fc" class="gj gk ct bk gl b dx gm gn r go" data-selectable-paragraph="">type CalculIncome struct {<br>  bonusRate int<br>}</span><span id="5103" class="gj gk ct bk gl b dx gp gq gr gs gt gn r go" data-selectable-paragraph="">func (c CalculIncome) VisitDeveloper(d Developer) {<br>  fmt.Println(d.Income + d.Income*c.bonusRate/100)<br>}</span><span id="aeec" class="gj gk ct bk gl b dx gp gq gr gs gt gn r go" data-selectable-paragraph="">func (c CalculIncome) VisitDirector(d Director) {<br>  fmt.Println(d.Income + d.Income*c.bonusRate/100)<br>}</span></pre><p id="4212" class="fj fk ct bk fl b fm fn fo fp fq fr fs ft fu fv fw cl" data-selectable-paragraph="">And I still need to inject it on our Employees, for that we need to add one modification to our Employee Interface, adding the Accept Method so Employee becomes :</p><pre class="gb gc gd ge gf gg gh gi"><span id="5093" class="gj gk ct bk gl b dx gm gn r go" data-selectable-paragraph="">type Employee interface {<br>  FullName()<br>  Accept(Visitor)<br>}</span></pre><p id="7207" class="fj fk ct bk fl b fm fn fo fp fq fr fs ft fu fv fw cl" data-selectable-paragraph="">And our objects</p><pre class="gb gc gd ge gf gg gh gi"><span id="4572" class="gj gk ct bk gl b dx gm gn r go" data-selectable-paragraph="">func (d Developer) Accept(v Visitor) {<br>  v.VisitDeveloper(d)<br>}</span><span id="b102" class="gj gk ct bk gl b dx gp gq gr gs gt gn r go" data-selectable-paragraph="">func (d Director) Accept(v Visitor) {<br>  v.VisitDirector(d)<br>}</span></pre><p id="1172" class="fj fk ct bk fl b fm fn fo fp fq fr fs ft fu fv fw cl" data-selectable-paragraph="">There we go, we have finished to implement our Visitor Pattern, let’s test it.</p><pre class="gb gc gd ge gf gg gh gi"><span id="df9d" class="gj gk ct bk gl b dx gm gn r go" data-selectable-paragraph="">backend := Developer{“Bob”, “Bilbo”, 1000, 32}<br>boss := Director{“Bob”, “Baggins”, 2000, 40}</span><span id="9bd4" class="gj gk ct bk gl b dx gp gq gr gs gt gn r go" data-selectable-paragraph="">backend.FullName()<br>backend.Accept(CalculIncome{20})</span><span id="7ceb" class="gj gk ct bk gl b dx gp gq gr gs gt gn r go" data-selectable-paragraph="">boss.FullName()<br>boss.Accept(CalculIncome{10})</span></pre><p id="dc21" class="fj fk ct bk fl b fm fn fo fp fq fr fs ft fu fv fw cl" data-selectable-paragraph="">output :</p><pre class="gb gc gd ge gf gg gh gi"><span id="63e1" class="gj gk ct bk gl b dx gm gn r go" data-selectable-paragraph="">Developer Bob Bilbo<br>1200</span><span id="db96" class="gj gk ct bk gl b dx gp gq gr gs gt gn r go" data-selectable-paragraph="">Director Bob Baggins<br>2200</span></pre><p id="ad2e" class="fj fk ct bk fl b fm fn fo fp fq fr fs ft fu fv fw cl" data-selectable-paragraph="">In the same wy, if we want to add more behaviors, we only need to implement visitor interface :</p><pre class="gb gc gd ge gf gg gh gi"><span id="372a" class="gj gk ct bk gl b dx gm gn r go" data-selectable-paragraph="">type AddingCaptainAge struct {<br>  captainAge int<br>}</span><span id="5b96" class="gj gk ct bk gl b dx gp gq gr gs gt gn r go" data-selectable-paragraph="">func (c AddingCaptainAge) VisitDeveloper(d Developer) {<br>  fmt.Println(d.Age + c.captainAge)<br>}</span><span id="6aa0" class="gj gk ct bk gl b dx gp gq gr gs gt gn r go" data-selectable-paragraph="">func (c AddingCaptainAge) VisitDirector(d Director) {<br>  fmt.Println(d.Age + c.captainAge)<br>}</span></pre><p id="42b5" class="fj fk ct bk fl b fm fn fo fp fq fr fs ft fu fv fw cl" data-selectable-paragraph="">and using it the same way as previsouly :</p><pre class="gb gc gd ge gf gg gh gi"><span id="964c" class="gj gk ct bk gl b dx gm gn r go" data-selectable-paragraph="">backend := Developer{“Bob”, “Bilbo”, 1000, 32}<br>boss := Director{“Bob”, “Baggins”, 2000, 40}</span><span id="426a" class="gj gk ct bk gl b dx gp gq gr gs gt gn r go" data-selectable-paragraph="">backend.FullName()<br>backend.Accept(CalculIncome{20})<br>backend.Accept(AddingCaptainAge{42})</span><span id="2ff9" class="gj gk ct bk gl b dx gp gq gr gs gt gn r go" data-selectable-paragraph="">boss.FullName()<br>boss.Accept(CalculIncome{10})<br>boss.Accept(AddingCaptainAge{42})</span></pre><p id="65b9" class="fj fk ct bk fl b fm fn fo fp fq fr fs ft fu fv fw cl" data-selectable-paragraph="">output :</p><pre class="gb gc gd ge gf gg gh gi"><span id="14ac" class="gj gk ct bk gl b dx gm gn r go" data-selectable-paragraph="">Developer  Bob   Bilbo<br>1200<br>74</span><span id="4115" class="gj gk ct bk gl b dx gp gq gr gs gt gn r go" data-selectable-paragraph="">Director  Bob   Baggins<br>2200<br>82</span></pre><p id="6109" class="fj fk ct bk fl b fm fn fo fp fq fr fs ft fu fv fw cl" data-selectable-paragraph="">Execute the code : <a href="https://play.golang.org/p/ddouZJj2z2S" class="at ey gu gv gw gx" target="_blank" rel="noopener nofollow">https://play.golang.org/p/ddouZJj2z2S</a></p></div>`
	longHtmlParsed   = `Visitor Design Pattern in golangFelipe Dutra Tine e SilvaFollowDec 19, 2019 · 2 min readWikipedia says :In object-oriented programming and software engineering, the visitor design pattern is a way of separating an algorithm from an object structure on which it operates. A practical result of this separation is the ability to add new operations to existing object structures without modifying the structures. It is one way to follow the open/closed principle.An example being better than 100 words, lets take an example.So we have our Employee Developer and Directors that is a structure that works well and are critical.type Employee interface {  FullName()}type Developer struct {  FirstName string  LastName string  Income float32  Age int}func (d Developer) FullName() {  fmt.Println(“Developer “,d.FirstName,” “,d.LastName)}type Director struct {  FirstName string  LastName string  Income float32  Age int}func (d Director) FullName() {  fmt.Println(“Director “,d.FirstName,” “,d.LastName)}An for some raison we want to *extend* that structure. But as these objects are highly critical we don’t want to change it (not too much).So We will use for that the Visitor Pattern, in order to “inject” new operations to existing object structures without modifying the structures.type Visitor interface {  VisitDeveloper(d Developer)  VisitDirector(d Director)}Now each new “additional Operation” or Behavior can be coded in a separate structure.type CalculIncome struct {  bonusRate int}func (c CalculIncome) VisitDeveloper(d Developer) {  fmt.Println(d.Income + d.Income*c.bonusRate/100)}func (c CalculIncome) VisitDirector(d Director) {  fmt.Println(d.Income + d.Income*c.bonusRate/100)}And I still need to inject it on our Employees, for that we need to add one modification to our Employee Interface, adding the Accept Method so Employee becomes :type Employee interface {  FullName()  Accept(Visitor)}And our objectsfunc (d Developer) Accept(v Visitor) {  v.VisitDeveloper(d)}func (d Director) Accept(v Visitor) {  v.VisitDirector(d)}There we go, we have finished to implement our Visitor Pattern, let’s test it.backend := Developer{“Bob”, “Bilbo”, 1000, 32}boss := Director{“Bob”, “Baggins”, 2000, 40}backend.FullName()backend.Accept(CalculIncome{20})boss.FullName()boss.Accept(CalculIncome{10})output :Developer Bob Bilbo1200Director Bob Baggins2200In the same wy, if we want to add more behaviors, we only need to implement visitor interface :type AddingCaptainAge struct {  captainAge int}func (c AddingCaptainAge) VisitDeveloper(d Developer) {  fmt.Println(d.Age + c.captainAge)}func (c AddingCaptainAge) VisitDirector(d Director) {  fmt.Println(d.Age + c.captainAge)}and using it the same way as previsouly :backend := Developer{“Bob”, “Bilbo”, 1000, 32}boss := Director{“Bob”, “Baggins”, 2000, 40}backend.FullName()backend.Accept(CalculIncome{20})backend.Accept(AddingCaptainAge{42})boss.FullName()boss.Accept(CalculIncome{10})boss.Accept(AddingCaptainAge{42})output :Developer  Bob   Bilbo120074Director  Bob   Baggins220082Execute the code : https://play.golang.org/p/ddouZJj2z2S`
	shortHtmlElement = `<td class="played">Hello World</td>`
	shortHtmlParsed  = `Hello World`
	longTexElement   = `\begin{document}

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
	longTexParsed = "\n\n\n(n.) A sacred fruit. Also known as:\n\n\nHere is the prevalence of each synonym.\n\nred lemon & uncommon life & common\n\n"
	shortTexElement = `\subsection*{How to handle topicalization}
I'll just assume a tree structure like \ex{1} .`
	shortTexParsed = `I'll just assume a tree structure like `
	)

func Test_visitor_ParseHtml(t *testing.T) {
	type args struct {
		html string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name:       "TestParseHtmlLong",
			args:       args{longHtmlElement},
			wantResult: longHtmlParsed,
		},
		{
			name:       "TestParseHtmlShort",
			args:       args{shortHtmlElement},
			wantResult: shortHtmlParsed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewVisitor()
			html := html_doc.NewHtmlDoc(tt.args.html)
			gotResult := v.ParseHtml(html)
			assert.Equal(t, tt.wantResult, gotResult)
		})
	}
}

func Test_visitor_ParseTex(t *testing.T) {
	type args struct {
		tex string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name:       "TestParseTexLong",
			args:       args{longTexElement},
			wantResult: longTexParsed,
		},
		{
			name:       "TestParseTexShort",
			args:       args{shortTexElement},
			wantResult: shortTexParsed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := NewVisitor()
			tex := tex_doc.NewTexDoc(tt.args.tex)
			gotResult := v.ParseTex(tex)
			assert.Equal(t, tt.wantResult, gotResult)
		})
	}
}
