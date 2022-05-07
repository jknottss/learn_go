<!DOCTYPE html>
<html lang="en" data-theme="light">
<head>

<link rel="preconnect" href="https://www.googletagmanager.com">
<script >(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
  new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
  j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
  'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
  })(window,document,'script','dataLayer','GTM-W8MVQXG');</script>
  
<meta charset="utf-8">
<meta name="description" content="Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="theme-color" content="#00add8">
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Material+Icons">
<link rel="stylesheet" href="/css/styles.css">

  
  <script>(function(w,d,s,l,i){w[l]=w[l]||[];w[l].push({'gtm.start':
  new Date().getTime(),event:'gtm.js'});var f=d.getElementsByTagName(s)[0],
  j=d.createElement(s),dl=l!='dataLayer'?'&l='+l:'';j.async=true;j.src=
  'https://www.googletagmanager.com/gtm.js?id='+i+dl;f.parentNode.insertBefore(j,f);
  })(window,document,'script','dataLayer','GTM-W8MVQXG');</script>
  
<script src="/js/site.js"></script>
<title> - The Go Programming Language</title>
</head>
<body class="Site">
  
<noscript><iframe src="https://www.googletagmanager.com/ns.html?id=GTM-W8MVQXG"
  height="0" width="0" style="display:none;visibility:hidden"></iframe></noscript>
  


<header class="Site-header js-siteHeader">
  <div class="Header Header--dark">
    <nav class="Header-nav">
      <a href="/">
        <img
          class="js-headerLogo Header-logo"
          src="/images/go-logo-white.svg"
          alt="Go">
      </a>
      <div class="Header-rightContent">
        <ul class="Header-menu">
          <li class="Header-menuItem ">
            <a href="/solutions/">Why Go</a>
          </li>
          <li class="Header-menuItem ">
            <a href="/learn/">Get Started</a>
          </li>
          <li class="Header-menuItem ">
            <a href="/doc/">Docs</a>
          </li>
          <li class="Header-menuItem ">
            <a href="https://pkg.go.dev">Packages</a>
          </li>
          <li class="Header-menuItem ">
            <a href="/play/">Play</a>
          </li>
          <li class="Header-menuItem ">
            <a href="/blog/">Blog</a>
          </li>
        </ul>
        <button class="Header-navOpen js-headerMenuButton Header-navOpen--white" aria-label="Open navigation.">
        </button>
      </div>
    </nav>
    
  </div>
</header>
<aside class="NavigationDrawer js-header">
  <nav class="NavigationDrawer-nav">
    <div class="NavigationDrawer-header">
      <a href="/">
        <img class="NavigationDrawer-logo" src="/images/go-logo-blue.svg" alt="Go.">
      </a>
    </div>
    <ul class="NavigationDrawer-list">
        <li class="NavigationDrawer-listItem ">
          <a href="/solutions/">Why Go</a>
        </li>
        <li class="NavigationDrawer-listItem ">
          <a href="/learn/">Get Started</a>
        </li>
        <li class="NavigationDrawer-listItem ">
          <a href="/doc/">Docs</a>
        </li>
        <li class="NavigationDrawer-listItem ">
          <a href="https://pkg.go.dev">Packages</a>
        </li>
        <li class="NavigationDrawer-listItem ">
          <a href="/play/">Play</a>
        </li>
        <li class="NavigationDrawer-listItem ">
          <a href="/blog/">Blog</a>
        </li>
    </ul>
  </nav>
</aside>
<div class="NavigationDrawer-scrim js-scrim" role="presentation"></div>
<main class="SiteContent SiteContent--default">
  

<article class="Texthtml Article">


<h1>Source file 


<a href="/src/">src</a>/<a href="/src/go/">go</a>/<a href="/src/go/ast/">ast</a>/<span class="text-muted">example_test.go</span>
</h1>


<pre><span id="L1" class="ln">     1&nbsp;&nbsp;</span><span class="comment">// Copyright 2012 The Go Authors. All rights reserved.</span>
<span id="L2" class="ln">     2&nbsp;&nbsp;</span><span class="comment">// Use of this source code is governed by a BSD-style</span>
<span id="L3" class="ln">     3&nbsp;&nbsp;</span><span class="comment">// license that can be found in the LICENSE file.</span>
<span id="L4" class="ln">     4&nbsp;&nbsp;</span>
<span id="L5" class="ln">     5&nbsp;&nbsp;</span>package ast_test
<span id="L6" class="ln">     6&nbsp;&nbsp;</span>
<span id="L7" class="ln">     7&nbsp;&nbsp;</span>import (
<span id="L8" class="ln">     8&nbsp;&nbsp;</span>	&#34;bytes&#34;
<span id="L9" class="ln">     9&nbsp;&nbsp;</span>	&#34;fmt&#34;
<span id="L10" class="ln">    10&nbsp;&nbsp;</span>	&#34;go/ast&#34;
<span id="L11" class="ln">    11&nbsp;&nbsp;</span>	&#34;go/format&#34;
<span id="L12" class="ln">    12&nbsp;&nbsp;</span>	&#34;go/parser&#34;
<span id="L13" class="ln">    13&nbsp;&nbsp;</span>	&#34;go/token&#34;
<span id="L14" class="ln">    14&nbsp;&nbsp;</span>)
<span id="L15" class="ln">    15&nbsp;&nbsp;</span>
<span id="L16" class="ln">    16&nbsp;&nbsp;</span><span class="comment">// This example demonstrates how to inspect the AST of a Go program.</span>
<span id="L17" class="ln">    17&nbsp;&nbsp;</span>func ExampleInspect() {
<span id="L18" class="ln">    18&nbsp;&nbsp;</span>	<span class="comment">// src is the input for which we want to inspect the AST.</span>
<span id="L19" class="ln">    19&nbsp;&nbsp;</span>	src := `
<span id="L20" class="ln">    20&nbsp;&nbsp;</span>package p
<span id="L21" class="ln">    21&nbsp;&nbsp;</span>const c = 1.0
<span id="L22" class="ln">    22&nbsp;&nbsp;</span>var X = f(3.14)*2 + c
<span id="L23" class="ln">    23&nbsp;&nbsp;</span>`
<span id="L24" class="ln">    24&nbsp;&nbsp;</span>
<span id="L25" class="ln">    25&nbsp;&nbsp;</span>	<span class="comment">// Create the AST by parsing src.</span>
<span id="L26" class="ln">    26&nbsp;&nbsp;</span>	fset := token.NewFileSet() <span class="comment">// positions are relative to fset</span>
<span id="L27" class="ln">    27&nbsp;&nbsp;</span>	f, err := parser.ParseFile(fset, &#34;src.go&#34;, src, 0)
<span id="L28" class="ln">    28&nbsp;&nbsp;</span>	if err != nil {
<span id="L29" class="ln">    29&nbsp;&nbsp;</span>		panic(err)
<span id="L30" class="ln">    30&nbsp;&nbsp;</span>	}
<span id="L31" class="ln">    31&nbsp;&nbsp;</span>
<span id="L32" class="ln">    32&nbsp;&nbsp;</span>	<span class="comment">// Inspect the AST and print all identifiers and literals.</span>
<span id="L33" class="ln">    33&nbsp;&nbsp;</span>	ast.Inspect(f, func(n ast.Node) bool {
<span id="L34" class="ln">    34&nbsp;&nbsp;</span>		var s string
<span id="L35" class="ln">    35&nbsp;&nbsp;</span>		switch x := n.(type) {
<span id="L36" class="ln">    36&nbsp;&nbsp;</span>		case *ast.BasicLit:
<span id="L37" class="ln">    37&nbsp;&nbsp;</span>			s = x.Value
<span id="L38" class="ln">    38&nbsp;&nbsp;</span>		case *ast.Ident:
<span id="L39" class="ln">    39&nbsp;&nbsp;</span>			s = x.Name
<span id="L40" class="ln">    40&nbsp;&nbsp;</span>		}
<span id="L41" class="ln">    41&nbsp;&nbsp;</span>		if s != &#34;&#34; {
<span id="L42" class="ln">    42&nbsp;&nbsp;</span>			fmt.Printf(&#34;%s:\t%s\n&#34;, fset.Position(n.Pos()), s)
<span id="L43" class="ln">    43&nbsp;&nbsp;</span>		}
<span id="L44" class="ln">    44&nbsp;&nbsp;</span>		return true
<span id="L45" class="ln">    45&nbsp;&nbsp;</span>	})
<span id="L46" class="ln">    46&nbsp;&nbsp;</span>
<span id="L47" class="ln">    47&nbsp;&nbsp;</span>	<span class="comment">// Output:</span>
<span id="L48" class="ln">    48&nbsp;&nbsp;</span>	<span class="comment">// src.go:2:9:	p</span>
<span id="L49" class="ln">    49&nbsp;&nbsp;</span>	<span class="comment">// src.go:3:7:	c</span>
<span id="L50" class="ln">    50&nbsp;&nbsp;</span>	<span class="comment">// src.go:3:11:	1.0</span>
<span id="L51" class="ln">    51&nbsp;&nbsp;</span>	<span class="comment">// src.go:4:5:	X</span>
<span id="L52" class="ln">    52&nbsp;&nbsp;</span>	<span class="comment">// src.go:4:9:	f</span>
<span id="L53" class="ln">    53&nbsp;&nbsp;</span>	<span class="comment">// src.go:4:11:	3.14</span>
<span id="L54" class="ln">    54&nbsp;&nbsp;</span>	<span class="comment">// src.go:4:17:	2</span>
<span id="L55" class="ln">    55&nbsp;&nbsp;</span>	<span class="comment">// src.go:4:21:	c</span>
<span id="L56" class="ln">    56&nbsp;&nbsp;</span>}
<span id="L57" class="ln">    57&nbsp;&nbsp;</span>
<span id="L58" class="ln">    58&nbsp;&nbsp;</span><span class="comment">// This example shows what an AST looks like when printed for debugging.</span>
<span id="L59" class="ln">    59&nbsp;&nbsp;</span>func ExamplePrint() {
<span id="L60" class="ln">    60&nbsp;&nbsp;</span>	<span class="comment">// src is the input for which we want to print the AST.</span>
<span id="L61" class="ln">    61&nbsp;&nbsp;</span>	src := `
<span id="L62" class="ln">    62&nbsp;&nbsp;</span>package main
<span id="L63" class="ln">    63&nbsp;&nbsp;</span>func main() {
<span id="L64" class="ln">    64&nbsp;&nbsp;</span>	println(&#34;Hello, World!&#34;)
<span id="L65" class="ln">    65&nbsp;&nbsp;</span>}
<span id="L66" class="ln">    66&nbsp;&nbsp;</span>`
<span id="L67" class="ln">    67&nbsp;&nbsp;</span>
<span id="L68" class="ln">    68&nbsp;&nbsp;</span>	<span class="comment">// Create the AST by parsing src.</span>
<span id="L69" class="ln">    69&nbsp;&nbsp;</span>	fset := token.NewFileSet() <span class="comment">// positions are relative to fset</span>
<span id="L70" class="ln">    70&nbsp;&nbsp;</span>	f, err := parser.ParseFile(fset, &#34;&#34;, src, 0)
<span id="L71" class="ln">    71&nbsp;&nbsp;</span>	if err != nil {
<span id="L72" class="ln">    72&nbsp;&nbsp;</span>		panic(err)
<span id="L73" class="ln">    73&nbsp;&nbsp;</span>	}
<span id="L74" class="ln">    74&nbsp;&nbsp;</span>
<span id="L75" class="ln">    75&nbsp;&nbsp;</span>	<span class="comment">// Print the AST.</span>
<span id="L76" class="ln">    76&nbsp;&nbsp;</span>	ast.Print(fset, f)
<span id="L77" class="ln">    77&nbsp;&nbsp;</span>
<span id="L78" class="ln">    78&nbsp;&nbsp;</span>	<span class="comment">// Output:</span>
<span id="L79" class="ln">    79&nbsp;&nbsp;</span>	<span class="comment">//      0  *ast.File {</span>
<span id="L80" class="ln">    80&nbsp;&nbsp;</span>	<span class="comment">//      1  .  Package: 2:1</span>
<span id="L81" class="ln">    81&nbsp;&nbsp;</span>	<span class="comment">//      2  .  Name: *ast.Ident {</span>
<span id="L82" class="ln">    82&nbsp;&nbsp;</span>	<span class="comment">//      3  .  .  NamePos: 2:9</span>
<span id="L83" class="ln">    83&nbsp;&nbsp;</span>	<span class="comment">//      4  .  .  Name: &#34;main&#34;</span>
<span id="L84" class="ln">    84&nbsp;&nbsp;</span>	<span class="comment">//      5  .  }</span>
<span id="L85" class="ln">    85&nbsp;&nbsp;</span>	<span class="comment">//      6  .  Decls: []ast.Decl (len = 1) {</span>
<span id="L86" class="ln">    86&nbsp;&nbsp;</span>	<span class="comment">//      7  .  .  0: *ast.FuncDecl {</span>
<span id="L87" class="ln">    87&nbsp;&nbsp;</span>	<span class="comment">//      8  .  .  .  Name: *ast.Ident {</span>
<span id="L88" class="ln">    88&nbsp;&nbsp;</span>	<span class="comment">//      9  .  .  .  .  NamePos: 3:6</span>
<span id="L89" class="ln">    89&nbsp;&nbsp;</span>	<span class="comment">//     10  .  .  .  .  Name: &#34;main&#34;</span>
<span id="L90" class="ln">    90&nbsp;&nbsp;</span>	<span class="comment">//     11  .  .  .  .  Obj: *ast.Object {</span>
<span id="L91" class="ln">    91&nbsp;&nbsp;</span>	<span class="comment">//     12  .  .  .  .  .  Kind: func</span>
<span id="L92" class="ln">    92&nbsp;&nbsp;</span>	<span class="comment">//     13  .  .  .  .  .  Name: &#34;main&#34;</span>
<span id="L93" class="ln">    93&nbsp;&nbsp;</span>	<span class="comment">//     14  .  .  .  .  .  Decl: *(obj @ 7)</span>
<span id="L94" class="ln">    94&nbsp;&nbsp;</span>	<span class="comment">//     15  .  .  .  .  }</span>
<span id="L95" class="ln">    95&nbsp;&nbsp;</span>	<span class="comment">//     16  .  .  .  }</span>
<span id="L96" class="ln">    96&nbsp;&nbsp;</span>	<span class="comment">//     17  .  .  .  Type: *ast.FuncType {</span>
<span id="L97" class="ln">    97&nbsp;&nbsp;</span>	<span class="comment">//     18  .  .  .  .  Func: 3:1</span>
<span id="L98" class="ln">    98&nbsp;&nbsp;</span>	<span class="comment">//     19  .  .  .  .  Params: *ast.FieldList {</span>
<span id="L99" class="ln">    99&nbsp;&nbsp;</span>	<span class="comment">//     20  .  .  .  .  .  Opening: 3:10</span>
<span id="L100" class="ln">   100&nbsp;&nbsp;</span>	<span class="comment">//     21  .  .  .  .  .  Closing: 3:11</span>
<span id="L101" class="ln">   101&nbsp;&nbsp;</span>	<span class="comment">//     22  .  .  .  .  }</span>
<span id="L102" class="ln">   102&nbsp;&nbsp;</span>	<span class="comment">//     23  .  .  .  }</span>
<span id="L103" class="ln">   103&nbsp;&nbsp;</span>	<span class="comment">//     24  .  .  .  Body: *ast.BlockStmt {</span>
<span id="L104" class="ln">   104&nbsp;&nbsp;</span>	<span class="comment">//     25  .  .  .  .  Lbrace: 3:13</span>
<span id="L105" class="ln">   105&nbsp;&nbsp;</span>	<span class="comment">//     26  .  .  .  .  List: []ast.Stmt (len = 1) {</span>
<span id="L106" class="ln">   106&nbsp;&nbsp;</span>	<span class="comment">//     27  .  .  .  .  .  0: *ast.ExprStmt {</span>
<span id="L107" class="ln">   107&nbsp;&nbsp;</span>	<span class="comment">//     28  .  .  .  .  .  .  X: *ast.CallExpr {</span>
<span id="L108" class="ln">   108&nbsp;&nbsp;</span>	<span class="comment">//     29  .  .  .  .  .  .  .  Fun: *ast.Ident {</span>
<span id="L109" class="ln">   109&nbsp;&nbsp;</span>	<span class="comment">//     30  .  .  .  .  .  .  .  .  NamePos: 4:2</span>
<span id="L110" class="ln">   110&nbsp;&nbsp;</span>	<span class="comment">//     31  .  .  .  .  .  .  .  .  Name: &#34;println&#34;</span>
<span id="L111" class="ln">   111&nbsp;&nbsp;</span>	<span class="comment">//     32  .  .  .  .  .  .  .  }</span>
<span id="L112" class="ln">   112&nbsp;&nbsp;</span>	<span class="comment">//     33  .  .  .  .  .  .  .  Lparen: 4:9</span>
<span id="L113" class="ln">   113&nbsp;&nbsp;</span>	<span class="comment">//     34  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {</span>
<span id="L114" class="ln">   114&nbsp;&nbsp;</span>	<span class="comment">//     35  .  .  .  .  .  .  .  .  0: *ast.BasicLit {</span>
<span id="L115" class="ln">   115&nbsp;&nbsp;</span>	<span class="comment">//     36  .  .  .  .  .  .  .  .  .  ValuePos: 4:10</span>
<span id="L116" class="ln">   116&nbsp;&nbsp;</span>	<span class="comment">//     37  .  .  .  .  .  .  .  .  .  Kind: STRING</span>
<span id="L117" class="ln">   117&nbsp;&nbsp;</span>	<span class="comment">//     38  .  .  .  .  .  .  .  .  .  Value: &#34;\&#34;Hello, World!\&#34;&#34;</span>
<span id="L118" class="ln">   118&nbsp;&nbsp;</span>	<span class="comment">//     39  .  .  .  .  .  .  .  .  }</span>
<span id="L119" class="ln">   119&nbsp;&nbsp;</span>	<span class="comment">//     40  .  .  .  .  .  .  .  }</span>
<span id="L120" class="ln">   120&nbsp;&nbsp;</span>	<span class="comment">//     41  .  .  .  .  .  .  .  Ellipsis: -</span>
<span id="L121" class="ln">   121&nbsp;&nbsp;</span>	<span class="comment">//     42  .  .  .  .  .  .  .  Rparen: 4:25</span>
<span id="L122" class="ln">   122&nbsp;&nbsp;</span>	<span class="comment">//     43  .  .  .  .  .  .  }</span>
<span id="L123" class="ln">   123&nbsp;&nbsp;</span>	<span class="comment">//     44  .  .  .  .  .  }</span>
<span id="L124" class="ln">   124&nbsp;&nbsp;</span>	<span class="comment">//     45  .  .  .  .  }</span>
<span id="L125" class="ln">   125&nbsp;&nbsp;</span>	<span class="comment">//     46  .  .  .  .  Rbrace: 5:1</span>
<span id="L126" class="ln">   126&nbsp;&nbsp;</span>	<span class="comment">//     47  .  .  .  }</span>
<span id="L127" class="ln">   127&nbsp;&nbsp;</span>	<span class="comment">//     48  .  .  }</span>
<span id="L128" class="ln">   128&nbsp;&nbsp;</span>	<span class="comment">//     49  .  }</span>
<span id="L129" class="ln">   129&nbsp;&nbsp;</span>	<span class="comment">//     50  .  Scope: *ast.Scope {</span>
<span id="L130" class="ln">   130&nbsp;&nbsp;</span>	<span class="comment">//     51  .  .  Objects: map[string]*ast.Object (len = 1) {</span>
<span id="L131" class="ln">   131&nbsp;&nbsp;</span>	<span class="comment">//     52  .  .  .  &#34;main&#34;: *(obj @ 11)</span>
<span id="L132" class="ln">   132&nbsp;&nbsp;</span>	<span class="comment">//     53  .  .  }</span>
<span id="L133" class="ln">   133&nbsp;&nbsp;</span>	<span class="comment">//     54  .  }</span>
<span id="L134" class="ln">   134&nbsp;&nbsp;</span>	<span class="comment">//     55  .  Unresolved: []*ast.Ident (len = 1) {</span>
<span id="L135" class="ln">   135&nbsp;&nbsp;</span>	<span class="comment">//     56  .  .  0: *(obj @ 29)</span>
<span id="L136" class="ln">   136&nbsp;&nbsp;</span>	<span class="comment">//     57  .  }</span>
<span id="L137" class="ln">   137&nbsp;&nbsp;</span>	<span class="comment">//     58  }</span>
<span id="L138" class="ln">   138&nbsp;&nbsp;</span>}
<span id="L139" class="ln">   139&nbsp;&nbsp;</span>
<span id="L140" class="ln">   140&nbsp;&nbsp;</span><span class="comment">// This example illustrates how to remove a variable declaration</span>
<span id="L141" class="ln">   141&nbsp;&nbsp;</span><span class="comment">// in a Go program while maintaining correct comment association</span>
<span id="L142" class="ln">   142&nbsp;&nbsp;</span><span class="comment">// using an ast.CommentMap.</span>
<span id="L143" class="ln">   143&nbsp;&nbsp;</span>func ExampleCommentMap() {
<span id="L144" class="ln">   144&nbsp;&nbsp;</span>	<span class="comment">// src is the input for which we create the AST that we</span>
<span id="L145" class="ln">   145&nbsp;&nbsp;</span>	<span class="comment">// are going to manipulate.</span>
<span id="L146" class="ln">   146&nbsp;&nbsp;</span>	src := `
<span id="L147" class="ln">   147&nbsp;&nbsp;</span>// This is the package comment.
<span id="L148" class="ln">   148&nbsp;&nbsp;</span>package main
<span id="L149" class="ln">   149&nbsp;&nbsp;</span>
<span id="L150" class="ln">   150&nbsp;&nbsp;</span>// This comment is associated with the hello constant.
<span id="L151" class="ln">   151&nbsp;&nbsp;</span>const hello = &#34;Hello, World!&#34; // line comment 1
<span id="L152" class="ln">   152&nbsp;&nbsp;</span>
<span id="L153" class="ln">   153&nbsp;&nbsp;</span>// This comment is associated with the foo variable.
<span id="L154" class="ln">   154&nbsp;&nbsp;</span>var foo = hello // line comment 2
<span id="L155" class="ln">   155&nbsp;&nbsp;</span>
<span id="L156" class="ln">   156&nbsp;&nbsp;</span>// This comment is associated with the main function.
<span id="L157" class="ln">   157&nbsp;&nbsp;</span>func main() {
<span id="L158" class="ln">   158&nbsp;&nbsp;</span>	fmt.Println(hello) // line comment 3
<span id="L159" class="ln">   159&nbsp;&nbsp;</span>}
<span id="L160" class="ln">   160&nbsp;&nbsp;</span>`
<span id="L161" class="ln">   161&nbsp;&nbsp;</span>
<span id="L162" class="ln">   162&nbsp;&nbsp;</span>	<span class="comment">// Create the AST by parsing src.</span>
<span id="L163" class="ln">   163&nbsp;&nbsp;</span>	fset := token.NewFileSet() <span class="comment">// positions are relative to fset</span>
<span id="L164" class="ln">   164&nbsp;&nbsp;</span>	f, err := parser.ParseFile(fset, &#34;src.go&#34;, src, parser.ParseComments)
<span id="L165" class="ln">   165&nbsp;&nbsp;</span>	if err != nil {
<span id="L166" class="ln">   166&nbsp;&nbsp;</span>		panic(err)
<span id="L167" class="ln">   167&nbsp;&nbsp;</span>	}
<span id="L168" class="ln">   168&nbsp;&nbsp;</span>
<span id="L169" class="ln">   169&nbsp;&nbsp;</span>	<span class="comment">// Create an ast.CommentMap from the ast.File&#39;s comments.</span>
<span id="L170" class="ln">   170&nbsp;&nbsp;</span>	<span class="comment">// This helps keeping the association between comments</span>
<span id="L171" class="ln">   171&nbsp;&nbsp;</span>	<span class="comment">// and AST nodes.</span>
<span id="L172" class="ln">   172&nbsp;&nbsp;</span>	cmap := ast.NewCommentMap(fset, f, f.Comments)
<span id="L173" class="ln">   173&nbsp;&nbsp;</span>
<span id="L174" class="ln">   174&nbsp;&nbsp;</span>	<span class="comment">// Remove the first variable declaration from the list of declarations.</span>
<span id="L175" class="ln">   175&nbsp;&nbsp;</span>	for i, decl := range f.Decls {
<span id="L176" class="ln">   176&nbsp;&nbsp;</span>		if gen, ok := decl.(*ast.GenDecl); ok &amp;&amp; gen.Tok == token.VAR {
<span id="L177" class="ln">   177&nbsp;&nbsp;</span>			copy(f.Decls[i:], f.Decls[i+1:])
<span id="L178" class="ln">   178&nbsp;&nbsp;</span>			f.Decls = f.Decls[:len(f.Decls)-1]
<span id="L179" class="ln">   179&nbsp;&nbsp;</span>			break
<span id="L180" class="ln">   180&nbsp;&nbsp;</span>		}
<span id="L181" class="ln">   181&nbsp;&nbsp;</span>	}
<span id="L182" class="ln">   182&nbsp;&nbsp;</span>
<span id="L183" class="ln">   183&nbsp;&nbsp;</span>	<span class="comment">// Use the comment map to filter comments that don&#39;t belong anymore</span>
<span id="L184" class="ln">   184&nbsp;&nbsp;</span>	<span class="comment">// (the comments associated with the variable declaration), and create</span>
<span id="L185" class="ln">   185&nbsp;&nbsp;</span>	<span class="comment">// the new comments list.</span>
<span id="L186" class="ln">   186&nbsp;&nbsp;</span>	f.Comments = cmap.Filter(f).Comments()
<span id="L187" class="ln">   187&nbsp;&nbsp;</span>
<span id="L188" class="ln">   188&nbsp;&nbsp;</span>	<span class="comment">// Print the modified AST.</span>
<span id="L189" class="ln">   189&nbsp;&nbsp;</span>	var buf bytes.Buffer
<span id="L190" class="ln">   190&nbsp;&nbsp;</span>	if err := format.Node(&amp;buf, fset, f); err != nil {
<span id="L191" class="ln">   191&nbsp;&nbsp;</span>		panic(err)
<span id="L192" class="ln">   192&nbsp;&nbsp;</span>	}
<span id="L193" class="ln">   193&nbsp;&nbsp;</span>	fmt.Printf(&#34;%s&#34;, buf.Bytes())
<span id="L194" class="ln">   194&nbsp;&nbsp;</span>
<span id="L195" class="ln">   195&nbsp;&nbsp;</span>	<span class="comment">// Output:</span>
<span id="L196" class="ln">   196&nbsp;&nbsp;</span>	<span class="comment">// // This is the package comment.</span>
<span id="L197" class="ln">   197&nbsp;&nbsp;</span>	<span class="comment">// package main</span>
<span id="L198" class="ln">   198&nbsp;&nbsp;</span>	<span class="comment">//</span>
<span id="L199" class="ln">   199&nbsp;&nbsp;</span>	<span class="comment">// // This comment is associated with the hello constant.</span>
<span id="L200" class="ln">   200&nbsp;&nbsp;</span>	<span class="comment">// const hello = &#34;Hello, World!&#34; // line comment 1</span>
<span id="L201" class="ln">   201&nbsp;&nbsp;</span>	<span class="comment">//</span>
<span id="L202" class="ln">   202&nbsp;&nbsp;</span>	<span class="comment">// // This comment is associated with the main function.</span>
<span id="L203" class="ln">   203&nbsp;&nbsp;</span>	<span class="comment">// func main() {</span>
<span id="L204" class="ln">   204&nbsp;&nbsp;</span>	<span class="comment">// 	fmt.Println(hello) // line comment 3</span>
<span id="L205" class="ln">   205&nbsp;&nbsp;</span>	<span class="comment">// }</span>
<span id="L206" class="ln">   206&nbsp;&nbsp;</span>}
<span id="L207" class="ln">   207&nbsp;&nbsp;</span>
</pre><p><a href="/src/go/ast/example_test.go?m=text">View as plain text</a></p>

</article>

</main>
<footer class="Site-footer">
  <div class="Footer">
    <div class="Container">
      <div class="Footer-links">
          <div class="Footer-linkColumn">
            <a href="/solutions/" class="Footer-link Footer-link--primary">
              Why Go
            </a>
              <a href="/solutions/#use-cases" class="Footer-link">
                Use Cases
              </a>
              <a href="/solutions/#case-studies" class="Footer-link">
                Case Studies
              </a>
          </div>
          <div class="Footer-linkColumn">
            <a href="/learn/" class="Footer-link Footer-link--primary">
              Get Started
            </a>
              <a href="/play" class="Footer-link">
                Playground
              </a>
              <a href="/tour/" class="Footer-link">
                Tour
              </a>
              <a href="https://stackoverflow.com/questions/tagged/go?tab=Newest" class="Footer-link">
                Stack Overflow
              </a>
              <a href="/help/" class="Footer-link">
                Help
              </a>
          </div>
          <div class="Footer-linkColumn">
            <a href="https://pkg.go.dev" class="Footer-link Footer-link--primary">
              Packages
            </a>
              <a href="/pkg/" class="Footer-link">
                Standard Library
              </a>
          </div>
          <div class="Footer-linkColumn">
            <a href="/project" class="Footer-link Footer-link--primary">
              About
            </a>
              <a href="/dl/" class="Footer-link">
                Download
              </a>
              <a href="/blog/" class="Footer-link">
                Blog
              </a>
              <a href="https://github.com/golang/go/issues" class="Footer-link">
                Issue Tracker
              </a>
              <a href="/doc/devel/release" class="Footer-link">
                Release Notes
              </a>
              <a href="https://blog.golang.org/go-brand" class="Footer-link">
                Brand Guidelines
              </a>
              <a href="/conduct" class="Footer-link">
                Code of Conduct
              </a>
          </div>
          <div class="Footer-linkColumn">
            <a href="https://www.twitter.com/golang" class="Footer-link Footer-link--primary">
              Connect
            </a>
              <a href="https://www.twitter.com/golang" class="Footer-link">
                Twitter
              </a>
              <a href="https://github.com/golang" class="Footer-link">
                GitHub
              </a>
              <a href="https://invite.slack.golangbridge.org/" class="Footer-link">
                Slack
              </a>
              <a href="https://reddit.com/r/golang" class="Footer-link">
                r/golang
              </a>
              <a href="https://www.meetup.com/pro/go" class="Footer-link">
                Meetup
              </a>
              <a href="https://golangweekly.com/" class="Footer-link">
                Golang Weekly
              </a>
          </div>
      </div>
    </div>
  </div>
  <div class="Footer">
    <div class="Container Container--fullBleed">
      <div class="Footer-bottom">
        <img class="Footer-gopher" src="/images/gophers/pilot-bust.svg" alt="The Go Gopher">
        <ul class="Footer-listRow">
          <li class="Footer-listItem">
            <a href="/copyright">Copyright</a>
          </li>
          <li class="Footer-listItem">
            <a href="/tos">Terms of Service</a>
          </li>
          <li class="Footer-listItem">
            <a href="http://www.google.com/intl/en/policies/privacy/"
              target="_blank"
              rel="noopener">
              Privacy Policy
            </a>
            </li>
          <li class="Footer-listItem">
            <a
              href="/s/website-issue"
              target="_blank"
              rel="noopener"
              >
              Report an Issue
            </a>
          </li>
        </ul>
        <a class="Footer-googleLogo" target="_blank" href="https://google.com" rel="noopener">
          <img class="Footer-googleLogoImg" src="/images/google-white.png" alt="Google logo">
        </a>
      </div>
    </div>
  </div>
  <script src="/js/jquery.js"></script>
  <script src="/js/carousels.js"></script>
  <script src="/js/searchBox.js"></script>
  <script src="/js/misc.js"></script>
  <script src="/js/hats.js"></script>
  <script src="/js/playground.js"></script>
  <script src="/js/godocs.js"></script>
</footer>
</body>
</html>














