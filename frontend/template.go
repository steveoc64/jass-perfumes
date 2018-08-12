package main

// This package has been automatically generated with temple.
// Do not edit manually!

import (
	"github.com/go-humble/temple/temple"
)

var (
	GetTemplate     func(name string) (*temple.Template, error)
	GetPartial      func(name string) (*temple.Partial, error)
	GetLayout       func(name string) (*temple.Layout, error)
	MustGetTemplate func(name string) *temple.Template
	MustGetPartial  func(name string) *temple.Partial
	MustGetLayout   func(name string) *temple.Layout
)

func init() {
	var err error
	g := temple.NewGroup()

	if err = g.AddTemplate("about", `<div class="about-content">
	
<img src="/img/flowerGS.png" class="constrainer">
<!-- <img src="img/chamelee.png"> -->
<span>
	EAU DE PARFUM
</span>
<p>
I’m barely human. I’m more like a creature; to me, everything gives off a scent!
<br>
Thoughts, moments, feelings, movements, words left unsaid, words barely spoken; they all have a distinct sense, distinct fragrances!
<br>
Both a smell and a touch!
<br>
To inhale is to capture, to experience!
<br>
I can perceive and I can “touch” in so many odd ways!
<br>
And so I am made up of all these scents, all these feelings!
<br>
An illumination of nerve endings.
<br>
I am Chamelee'.
</p>

<div class="blog-shares">
		<div>
				<i class="share-twitter fa fa-twitter-square fa-3x"></i>
		</div>

		<div class="fb-like" data-href="{{.GetURL}}" data-layout="standard" data-action="like" data-size="small" data-show-faces="true" data-share="true">
		<i class="share-facebook fa fa-facebook-square fa-3x"></i>	
		</div>

		<div class="g-plusone" data-width="300" data-href="{{.GetURL}}">
			<a href="https://plus.google.com/share?url={{.GetURL}}">
			<i class="share-google fa fa-google-plus-square fa-3x"></i>	
			</a>
		</div>
</div>

<div class="jass-logo-small"> </div>
</div>`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("cart", `{{range $key,$value := .CartItems}}
<div>
	{{$value}}
</div>
{{end}}

<div class="jass-logo-small-box"> </div>`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("contact", `Chamelee Contact Form
`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("footer", `	<img src="/img/flowerGS.png">
	<span class="backbtn">
		Discover More ...
	</span>
	<span class="backbtn footer-x">
	<i class="fa fa-chevron-down"></i>	
	</span>

	<div class="backbtn">
		More
		<div data-href="/shop">Shop</div>
		<div data-href="/blog">Blog</div>
		<div data-href="/ambassadors">Ambassadors</div>
		<div data-href="/privacy">Terms & Cond.</div>
		<div data-href="/privacy">Privacy</div>
		<div data-href="/about">About</div>
		<hr>
		<div data-href="/newsletter">Newsletter</div>
		<div data-href="/contact">Contact Us</div>
	</div>
	<div class="backbtn social-column">
		Follow us
		<i class="fa fa-instagram fa-2x"></i>
		<i class="fa fa-facebook-square fa-2x"></i>	
		<i class="fa fa-twitter-square fa-2x"></i>
		<i class="fa fa-pinterest-square fa-2x"></i>
		<i class="fa fa-google-plus-square fa-2x"></i>	
		<i class="fa fa-youtube-square fa-2x"></i>	
	</div>`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("jass-blog-article", `<div class="blog-article-name gotop">
	<span class="gotop">{{.Name}}</span>
	<i class="fa fa-chevron-up"></i>
</div>
<div class="blog-article">
	<div class="blog-article-image">
		<div class="image-box">
			<img src="/img/models/{{.Image}}">
		</div>
	</div>
	<div class="blog-article-title">{{.Title}}</div>
	<div class="blog-article-content">{{range .GetLines}} {{.}}<br> {{end}}</div>

	<div class="blog-shares">
		<div>
			<a class="twitter-share-button" href="https://twitter.com/intent/tweet?text={{.GetURL}}+{{.Name}}" data-size="large">
				<i class="share-twitter fa fa-twitter-square fa-3x"></i>
			</a>
		</div>

		<div class="fb-like" data-href="{{.GetURL}}" data-layout="standard" data-action="like" data-size="small" data-show-faces="true" data-share="true">
		<i class="share-facebook fa fa-facebook-square fa-3x"></i>	
		</div>

		<div class="g-plusone" data-width="300" data-href="{{.GetURL}}">
			<a href="https://plus.google.com/share?url={{.GetURL}}">
			<i class="share-google fa fa-google-plus-square fa-3x"></i>	
			</a>
		</div>
	</div>

	<div class="jass-logo-small">
		<i class="fa fa-chevron-left fa-3x"></i>
		<i class="fa fa-chevron-right fa-3x"></i>
	</div>
</div>
`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("jass-blog", `<div class="blog-header"><i class="fa fa-angle-up"></i></div>
<div class="blog-container">
<div class="blog-item">
<iframe width="1040" height="630" src="https://www.youtube.com/embed/Daoy9EUKWu8" frameborder="0" allowfullscreen></iframe>
</div>
<div class="blog-item">
	<iframe width="1040" height="630" src="https://www.youtube.com/embed/GKNecMv9BdY" frameborder="0" allowfullscreen></iframe>
</div>
{{range $key,$value := .Blogs}}
<div class="blog-item" name="blog-{{$value.ID}}" data-id="{{$value.ID}}">
	<div class="blog-item-pic" name="blog-image-{{$value.ID}}" data-id="{{$value.ID}}"></div>
	<div class="blog-item-title" name="blog-title-{{$value.ID}}" data-id="{{$value.ID}}">
		{{$value.Name}}
	</div>
</div>
{{end}}
<div class="jass-logo-small"></div>
</div>
`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("main-page", `<div class="container">
</div>
`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("privacy", `<div class="about-content privacy-policy">
	
<img src="/img/flowerGS.png" class="constrainer">

<h4>
Chamelee Privacy Policy
</h4>

<p>
Chamelee has created this privacy statement in order to demonstrate our firm commitment to privacy.
</p>

<h4>
General Statement of Principles 
</h4>

<p>
As described in detail below, any information we gather at this web site is strictly for our use and is not shared with any other entity, public or private, for any reason - period. We will not sell or give away any lists or other data that we may retain and we do not purchase such information from other sources.
</p>

<h4>
Statistical Data 
</h4>

<p>
Our servers (as most) track IP addresses and referring pages to help with site maintenance and improvements. This data is viewed only as anonymous statistics to show the busiest times of the day or week, pages with errors and how effective our advertising has been. This information is not used for any other purpose.
</p>

<h4>
Personal Information Collected - Order Forms
</h4>

<p>
With the exception of credit card info, we store the information from your order form to allow us to track consulting issues or refer to a previous order to help provide some customer service. You may elect to have your information completely removed from this system by e-mailing us at info@chamelee.com.au with your request. We do not store any type of sales information.
</p>

<h4>
Information Correction or Removal 
</h4>

<p>
If you wish to correct, update or remove any information about you that may be in our records, please send us e-mail at info@chamelee.com.au with the details of your request. If you wish to contact us further, please find complete contact information on our contact page.
</p>

<div class="jass-logo-small"> </div>
</div>`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("sale-cats", `<div>
{{range $key,$value := .Category}}
	<span class="sale-category" data-id="{{$value.ID}}">{{$value.Name}}</span>
{{end}}
</div>`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("sale-items", `{{$catID := .SelectCatID}}
{{range $key,$value := .Products}}

{{if eq $catID 0}}
<div class="jass-sale-item" data-id="{{$value.ID}}">
	<img src="/img/product/{{$value.Image}}">
	<div class="sale-name">
		{{$value.Name}}
	</div>
 	<div class="sale-price-small">
 		$ {{$value.Price}}
 	</div>
</div>
{{else if eq $catID $value.CatID}}
<div class="jass-sale-item" data-id="{{$value.ID}}">
	<img src="/img/product/{{$value.Image}}">
	<div class="sale-name">
		{{$value.Name}}
	</div>
 	<div class="sale-price-small">
 		$ {{$value.Price}}
 	</div>
</div>
{{end}}

{{end}}


<div class="jass-logo-small"></div>`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("sales-bar", `<span class="sales-account"><i class="fa fa-2x fa-user"></i></span>
<span class="sales-qty">{{.GetCartItemCount}}</span>
<span class="sales-amount">{{.GetCartTotal}}</span>
<span class="sales-cart"><i class="fa fa-2x fa-shopping-cart"></i></span>
<!-- <span class="sales-checkout hidden"><i class="fa fa-2x fa-credit-card"></i></span> -->`); err != nil {
		panic(err)
	}

	if err = g.AddTemplate("slidemenu", `<!-- Slide in menu once logged in  -->
<nav class="cbp-spmenu cbp-spmenu-vertical cbp-spmenu-right" id="slidemenu">
  <a href="#" id="menu-fragrances"><i class="fa fa-snowflake-o"></i> Fragrances</a>
  <a href="#" id="menu-skincare"><i class="fa fa-hand-lizard-o"></i> Skincare</a>
  <a href="#" id="menu-merchandise"><i class="fa fa-gift"></i> Merchandise</a>
  <a href="#" id="menu-blog"><i class="fa fa-hashtag"></i> Blog</a>
  <a href="#" id="menu-ambassadors"><i class="fa fa-user-circle-o"></i> Ambassadors</a>
  <a href="#" id="menu-about"><i class="fa fa-question-circle-o"></i> About</a>
  <a href="#" id="menu-contact"><i class="fa fa-at"></i> Contact</a>
<!-- 
  <a href="#" id="menu-shop"><i class="fa fa-shopping-bag"></i> Shop</a>
  <a href="#" id="menu-discover"><i class="fa fa-snowflake-o"></i> Discover</a>
  <a href="#" id="menu-merchandise"><i class="fa fa-gift"></i> Merchandise</a>
  <a href="#" id="menu-facebook"><i class="fa fa-facebook"></i> Facebook</a>
  <a href="#" id="menu-twitter"><i class="fa fa-twitter"></i> Twitter</a>
  <a href="#" id="menu-instagram"><i class="fa fa-instagram"></i> Instagram</a>
  <a href="#" id="menu-blog"><i class="fa fa-hashtag"></i> Blog</a>
  <a href="#" id="menu-contact"><i class="fa fa-at"></i> Contact</a>

 --></nav> 



`); err != nil {
		panic(err)
	}

	GetTemplate = g.GetTemplate
	GetPartial = g.GetPartial
	GetLayout = g.GetLayout
	MustGetTemplate = g.MustGetTemplate
	MustGetPartial = g.MustGetPartial
	MustGetLayout = g.MustGetLayout
}
