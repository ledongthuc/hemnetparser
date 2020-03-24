package hemnetparser

import "testing"

var document string = `<!DOCTYPE html>
<html lang="sv-SE" dir="ltr">
  <!--
           hhhhhhhh
      hhhhhhhhhhhhhhhhhh
    hhhhh             hhhh
  hhhhh                 hhhh
 hhhh         hhh        hhhh
 hhh       hhhhhhhhh      hhhh
hhhh     hhhhhhhhhhhh     hhhh
hhhh     hhhh    hhhh     hhhh
hhhh     hhhh    hhhh     hhhh
 hhhh                    hhhh
  hhhhh                 hhhh
   hhhhhh            hhhhh
 hhhhh hhhhhhhhhhhhhhhh
hhhh     hhhhhhhhhhh

Kul att du hittat hit! Vill du också jobba på en av Sveriges största sajter med miljontals unika besökare?
Gör en spontanansökan på https://jobba.hemnet.se/
-->

  <head>
    <meta charset="utf-8">
    <link rel="dns-prefetch" href="//assets.hemnet.se">
    <link rel="dns-prefetch" href="//bilder.hemnet.se">
    <link rel="dns-prefetch" href="//maps.googleapis.com">

    <title>Jämtlandsgatan 152, 6 tr i Vällingby Centrum, Stockholm - Bostadsrätt till salu - Hemnet</title>

    <meta name="format-detection" content="telephone=no">
    <meta name="google-site-verification" content="MgOW5i2sIhcrdILu7eVDwzPcciV8KK3o3BEI2UNnqGQ" />
    <meta name="description" content="Etta om 33 kvm i populära brf Fornkullen nära T-bana och Vällingby C.  Möjlig att göra om till en tvåa. Balkong i västerläge med fantastiks utsikt. Lägenhete...">
    <meta name="apple-mobile-web-app-title" content="Hemnet">

    <meta name="viewport" content="initial-scale=1.0,width=device-width">

    <link rel="preload" href="https://assets.hemnet.se/assets/fonts/opensans-semibold-webfont.woff2" as="font" type="font/woff2" crossorigin>
    <link rel="preload" href="https://assets.hemnet.se/assets/fonts/opensans-regular-webfont.woff2" as="font" type="font/woff2" crossorigin>
    <link rel="preload" href="https://assets.hemnet.se/assets/fonts/hemnetslabbold_mrf1301-webfont.woff2" as="font" type="font/woff2" crossorigin>

    <meta name="csrf-param" content="authenticity_token" />
    <link rel="shortcut icon" type="image/png" href="/favicon.ico" />
    <link rel="apple-touch-icon" href="https://assets.hemnet.se/assets/images/apple-touch-icon.png">

    

    <link rel="stylesheet" media="screen" href="https://assets.hemnet.se/assets/packs/application.3f88d3c4e7fda6ff6e03.css" />
      <link rel="stylesheet" media="screen, print" href="https://assets.hemnet.se/assets/packs/listings.67228302a2d9846ac499.css" />
  <link rel="stylesheet" media="print" href="https://assets.hemnet.se/assets/packs/property-print.16914c83a9e6fa7adaf0.css" />
    <link rel="stylesheet" media="screen" href="https://assets.hemnet.se/assets/packs/housing-cooperative.673f05e376632898e8a4.css" />
  <link rel="stylesheet" media="screen" href="https://assets.hemnet.se/assets/packs/header-footer.396c4e4498794a66d0f7.css" />


    <script type="application/ld+json">
  {
    "@context": "http://schema.org",
    "@type": "WebSite",
    "name": "Hemnet",
    "url": "https://www.hemnet.se"
  }
  {
    "@context": "http://schema.org",
    "@type": "Organization",
    "url": "https://www.hemnet.se",
    "logo": "https://assets.hemnet.se/assets/images/hemnet-logo.svg"
  }
</script>

      

<link rel="canonical" href="https://www.hemnet.se/bostad/lagenhet-1rum-vallingby-centrum-stockholms-kommun-jamtlandsgatan-152,-6-tr-16663146">
<meta name="apple-itunes-app" content="app-id=525017304, app-argument=hemnet://listings/16663146">

  <meta property="og:title" content="Jämtlandsgatan 152, 6 tr, Vällingby Centrum Stockholms kommun">
<meta property="og:url" content="https://www.hemnet.se/bostad/lagenhet-1rum-vallingby-centrum-stockholms-kommun-jamtlandsgatan-152,-6-tr-16663146">
<meta property="og:site_name" content="Hemnet">
<meta property="og:description" content="Etta om 33 kvm i populära brf Fornkullen nära T-bana och Vällingby C.  Möjlig att göra om till en tvåa. Balkong i västerläge med fantastiks utsikt. Lägenheten målades om 2019.   Under 2018 har före...">
<meta property="og:type" content="website">

  <meta property="og:image" content="https://bilder.hemnet.se/images/itemgallery_cut/ca/16/ca162533055df0661ab81070ffe8d308.jpg">
  <meta property="og:image:width" content="690">
  <meta property="og:image:height" content="460">

  <meta name="twitter:card" content="summary_large_image">
<meta name="twitter:site" content="@hemnet">
<meta name="twitter:domain" content="https://www.hemnet.se">
<meta name="twitter:title" content="Jämtlandsgatan 152, 6 tr, Vällingby Centrum Stockholms kommun">
<meta name="twitter:description" content="Etta om 33 kvm i populära brf Fornkullen nära T-bana och Vällingby C.  Möjlig att göra om till en tvåa. Balkong i västerläge med fantastiks utsikt. Lägenheten målades om 2019.   Under 2018 har före...">

  <meta name="twitter:image:src" content="https://bilder.hemnet.se/images/itemgallery_cut/ca/16/ca162533055df0661ab81070ffe8d308.jpg">


    <script type="application/ld+json">
      {
        "@context": "http://schema.org",
        "@type": "Event",
        "name": "Visning Imorgon kl. 00:00 av Jämtlandsgatan 152, 6 tr",
        "description": "Välkommen på visning – för tider,mer info och anmälan gå till Mäklarhusets sida",
        "image": "https://bilder.hemnet.se/images/itemgallery_L/ca/16/ca162533055df0661ab81070ffe8d308.jpg",
        "startDate" : "2020-03-25T00:00:00+0100",
        "endDate": "2020-03-25T23:59:59+0100",
        "url" : "https://www.hemnet.se/bostad/lagenhet-1rum-vallingby-centrum-stockholms-kommun-jamtlandsgatan-152,-6-tr-16663146",
          "offers": {
            "@type": "Offer",
            "priceCurrency": "SEK",
            "price": 1875000,
            "priceValidUntil": "2022-03-09T17:29:55+0100",
            "availability": "http://schema.org/InStock",
            "validFrom": "2020-03-09T17:29:55+0100",
            "url" : "https://www.hemnet.se/bostad/lagenhet-1rum-vallingby-centrum-stockholms-kommun-jamtlandsgatan-152,-6-tr-16663146"
          },
          "performer" : {
            "@type": "Person",
            "name":  "Robert Blomster",
            "memberOf": {
              "@type": "Organization",
                "name": "Mäklarhuset Vällingby-Hässelby"
            }
          },
        "location" : {
          "@type" : "Place",
          "name" : "Lägenhet på Jämtlandsgatan 152, 6 tr",
          "address" : {
            "@type": "PostalAddress",
            "streetAddress" : "Jämtlandsgatan 152, 6 tr",
            "addressLocality": "Vällingby Centrum, Stockholms kommun",
            "addressRegion": "Stockholms län",
            "postalCode": 
          }
        }
      }
    </script>
    <script type="application/ld+json">
      {
        "@context": "http://schema.org",
        "@type": "Event",
        "name": "Visning Sön 29 mar kl. 00:00 av Jämtlandsgatan 152, 6 tr",
        "description": "Välkommen på visning – för tider,mer info och anmälan gå till Mäklarhusets sida",
        "image": "https://bilder.hemnet.se/images/itemgallery_L/ca/16/ca162533055df0661ab81070ffe8d308.jpg",
        "startDate" : "2020-03-29T00:00:00+0100",
        "endDate": "2020-03-29T23:59:59+0200",
        "url" : "https://www.hemnet.se/bostad/lagenhet-1rum-vallingby-centrum-stockholms-kommun-jamtlandsgatan-152,-6-tr-16663146",
          "offers": {
            "@type": "Offer",
            "priceCurrency": "SEK",
            "price": 1875000,
            "priceValidUntil": "2022-03-09T17:29:55+0100",
            "availability": "http://schema.org/InStock",
            "validFrom": "2020-03-09T17:29:55+0100",
            "url" : "https://www.hemnet.se/bostad/lagenhet-1rum-vallingby-centrum-stockholms-kommun-jamtlandsgatan-152,-6-tr-16663146"
          },
          "performer" : {
            "@type": "Person",
            "name":  "Robert Blomster",
            "memberOf": {
              "@type": "Organization",
                "name": "Mäklarhuset Vällingby-Hässelby"
            }
          },
        "location" : {
          "@type" : "Place",
          "name" : "Lägenhet på Jämtlandsgatan 152, 6 tr",
          "address" : {
            "@type": "PostalAddress",
            "streetAddress" : "Jämtlandsgatan 152, 6 tr",
            "addressLocality": "Vällingby Centrum, Stockholms kommun",
            "addressRegion": "Stockholms län",
            "postalCode": 
          }
        }
      }
    </script>

<script type="application/ld+json">
  {
    "@context": "http://schema.org/",
    "@type": "Product",
    "name": "Jämtlandsgatan 152, 6 tr",
    "image": "https://bilder.hemnet.se/images/itemgallery_L/ca/16/ca162533055df0661ab81070ffe8d308.jpg",
    "description":  "Etta om 33 kvm i populära brf Fornkullen nära T-bana och Vällingby C. Möjlig att göra om till en tvåa. Balkong i västerläge med fantastiks utsikt. Lägenheten målades om 2019. Under 2018 har föreningen bytt fönster och balkongdörrar. Förening med mycket bra ekonomi, låg månadsavgift. Lugnt och centralt läge. Hiss och tvättstuga i huset. St. Stambytt helkaklat badrum. Utmärkt läge endast ca 200 meter till Vällingby C med butiker, restauranger, biograf, T-bana/buss och gym. Nära till simhall, Grimsta Naturreservat med motionsspår, badplatser, minigolf och fiske vid Mälaren. Välkommen på visning!",
      "offers": {
        "@type": "Offer",
        "priceCurrency": "SEK",
        "price": 1875000,
        "priceValidUntil": "2022-03-09T17:29:55+0100",
        "availability": "http://schema.org/InStock",
        "validFrom": "2020-03-09T17:29:55+0100",
        "url": "https://www.hemnet.se/bostad/lagenhet-1rum-vallingby-centrum-stockholms-kommun-jamtlandsgatan-152,-6-tr-16663146"
      },
    "mpn": "16663146",
    "brand": "Mäklarhuset Vällingby-Hässelby"
  }
</script>

<script type="application/ld+json">
{
  "@context": "http://schema.org",
  "@type": "Place",
  "address" : {
    "@type": "PostalAddress",
    "streetAddress" : "Jämtlandsgatan 152, 6 tr",
    "addressLocality": "Vällingby Centrum, Stockholms kommun",
    "addressRegion": "Stockholms län",
    "addressCountry": "SE",
    "postalCode": 
  }
}
</script>



    <script src="https://assets.hemnet.se/assets/javascripts/legacy/h3_head.min.js?1585059633"></script>
    <script>
      Hemnet = Hemnet || {};
      Hemnet.server = Hemnet.server || {};
      Hemnet.server.data = {"csrf":{"value":"GKgjkGovwO8XKxkA2DibxA7BbcKl0aN+aLAOxpum8P+PsssH+qDec1QffZwmoz061WRWx386mM/+JD9ie8cujQ==","name":"authenticity_token"},"environment":"prod","isInternalTraffic":false,"refreshPath":"/open-api/server-data.json","sentryDSN":"https://d413eda448c64289bd3315951e24b44f@sentry.io/276700","search":{"municipalitiesPath":"/locations/municipalities","newSessionPath":"/mitt_hemnet/session/logga_in"}};
    </script>

    <script>
  var googletag = googletag || {};
  googletag.cmd = googletag.cmd || [];
</script>

<script src="https://securepubads.g.doubleclick.net/tag/js/gpt.js" async="async"></script>
  <script>
    googletag.cmd.push(function() {
      var slot = googletag.defineSlot('/20339512601/d_p_listing', [[980, 240], [980, 120], [1, 1]], "ad-unit-0");
      googletag.pubads().collapseEmptyDivs(false);

        slot.setTargeting("item_types", [7]);
        slot.setTargeting("rooms", [1]);
        slot.setTargeting("price", [2]);
        slot.setTargeting("municipality", [199]);
        slot.setTargeting("district", [109]);
        slot.setTargeting("region", [12]);
        slot.setTargeting("country", [1]);
          slot.setTargeting("property_price", 1875000);
          slot.setTargeting("property_address", "Jämtlandsgatan 152, 6 tr");
          slot.setTargeting("property_living_area", 33);
          slot.setTargeting("property_zipcode", "");
      slot.setTargeting("page", "listing");
      slot.setTargeting("position", "1");
      slot.addService(googletag.pubads());
      var mapping = googletag.sizeMapping()
        .addSize([740, 0], [[980, 240], [980, 120], [1, 1]])
        .addSize([0, 0], [])
        .build();
      slot.defineSizeMapping(mapping);

      if (true) {
        var placeholder = document.getElementById("ad_unit_placeholder");
        googletag.pubads().addEventListener('slotRenderEnded', function(event) {
          placeholder.remove();
        });
      }
    });
  </script>

<script async>
  function setupAds() {
    var windowWidth = Math.max(document.documentElement.clientWidth, window.innerWidth || 0);
    var isMobile = windowWidth <= 670;
    var isDesktop = !isMobile;
    var isStartPage = window.location.pathname === "/";

    googletag.cmd.push(function() {
      var hasConsentedToAds = function(name) {
        var value = localStorage.getItem(name) ? JSON.parse(localStorage.getItem(name)).value : ""
        var parts = value.split("+")
        return parts[0] === "YES"
      }

      var usePersonalizedAds = hasConsentedToAds("_hemnet_consent") ? 0 : 1;

      googletag.pubads().setRequestNonPersonalizedAds(usePersonalizedAds);

      if (isDesktop && !isStartPage) {
        googletag.pubads().enableSingleRequest();
      }

      googletag.pubads().setSafeFrameConfig({
        allowPushExpansion: true,
        allowOverlayExpansion: false,
        sandbox: true
      });
      googletag.pubads().enableLazyLoad({
        fetchMarginPercent: 2000, // Fetching ads 20 viewports below scroll position
        renderMarginPercent: 25,
        mobileScaling: 2.0
      });
      googletag.pubads().addEventListener('slotRenderEnded', function slotRendered(e) {
        var slotEl;
        var adLabelEl;
        var isInhouseCampaign = e.advertiserId === 65047801;
        var isOrderedInhouseCampaign = e.campaignId === 2546920719;

        if (!e.isEmpty) {
          slotEl = document.getElementById(e.slot.getSlotElementId());
          adLabelEl = document.createElement('a');

          if (isOrderedInhouseCampaign || !isInhouseCampaign) {
            adLabelEl.href = '/om/cookies';
            adLabelEl.innerHTML = 'annons';
            adLabelEl.className = 'ad-unit__ad-label';
          }

          slotEl.insertBefore(adLabelEl, slotEl.firstChild);
          var isWallpaper = e.size.indexOf(1440) >= 0;
          var noAdsPixel = e.size.indexOf(1) >= 0;
          var is980 = e.size.indexOf(980) == 0;

          if (noAdsPixel) {
            slotEl.classList.add('ad-unit--no-ad');
          }
          if (isWallpaper) {
            slotEl.classList.add('ad-unit--top-wallpaper');
            document.body.classList.add('wallpaper-ad');
            nagbar = document.getElementsByClassName("nagbar")[0];

            if (nagbar !== undefined) {
              slotEl.classList.add('ad-unit--top-wallpaper__with-nagbar');
            }
          }
          if (is980) {
            slotEl.classList.add('ad-unit--980');
          }
        }
      });
      
      googletag.enableServices();
    });
  }

  if (document.attachEvent ? document.readyState === 'complete' : document.readyState !== 'loading') {
    window.setTimeout(setupAds(), 0);
  } else {
    document.addEventListener('DOMContentLoaded', setupAds());
  }
</script>


    
  </head>
  <body>
    <div class="hemnet-modal-portal"></div>

      
<div class="site-header qa-site-header">
  <div class="site-header__logo">
    
<a onclick="hnt.ga(&#39;toppmeny&#39;, &#39;logo&#39;, this.href);" href="/">
  <img onerror="this.src=&quot;https://assets.hemnet.se/assets/images/hemnet-symbol.png&quot;;this.onerror=null;" alt="Hemnet" class="logo-small" src="https://assets.hemnet.se/assets/images/hemnet-symbol.svg" width="28" height="28" />
  <img onerror="this.src=&quot;https://assets.hemnet.se/assets/images/hemnet-logo.png&quot;;this.onerror=null;" alt="Hemnet" class="logo-large" src="https://assets.hemnet.se/assets/images/hemnet-logo.svg" width="120" height="28" />
</a>
  </div>

  <div class="site-header__controls">
      <div class="site-header__control">
        <a class="site-header__search-button" onclick="hnt.ga(&#39;toppmeny&#39;, &#39;sök-knapp&#39;, this.href);" href="/">
          <i class="fa fa-search"></i> Sök
</a>      </div>

    <div class="site-header__control">
      <button class="site-header__menu-button js-primary-nav-toggle" onclick="hnt.ga(&#39;toppmeny&#39;, &#39;expandering responsiv meny&#39;);">
        <i class="fa fa-bars"></i> Mer
      </button>
    </div>
  </div>

  <div class="site-header__primary-nav">
    
<nav class="primary-nav js-primary-nav">

  <ul class="primary-nav__list">
    <li class="primary-nav__list-item">
      <a href="/"
         class="js-primary-nav-dropdown-anchor qa-primary-nav-find-property primary-nav__link"
         onclick="hnt.ga(&#39;toppmeny&#39;, &#39;sök bostad&#39;, &#39;Utfälld meny&#39;);">
        Sök bostad
        <i class="primary-nav__icon fa fa-angle-down"></i>
      </a>

      <div class="primary-nav-dropdown primary-nav-dropdown--find">
        <ul class="primary-nav-dropdown__link-list">
          <li class="primary-nav-dropdown__link-list-item">
            <a class="qa-primary-nav-find-property-new primary-nav-find-property__link" onclick="hnt.ga(&#39;toppmeny&#39;, &#39;sök bostad&#39;, &#39;Ny sökning&#39;);" href="/">Ny sökning</a>
          </li>
          <li class="primary-nav-dropdown__link-list-item">
            <a class="qa-primary-nav-find-property-all primary-nav-find-property__link" onclick="hnt.ga(&#39;toppmeny&#39;, &#39;sök bostad&#39;, &#39;Alla bostäder&#39;);" href="/sitemap">Alla bostäder</a>
          </li>
        </ul>

        <div class="primary-nav-find-property__search-history">
          <div class="primary-nav-find-property__heading">Senaste sökningar</div>
          <div class="js-navigation-history"></div>
        </div>
      </div>
    </li>
      <li class="primary-nav__list-item">
        <a href="/salja-bostad"
          class="js-primary-nav-dropdown-anchor qa-primary-nav-sell-property primary-nav__link primary-nav__link--sell"
          onclick="hnt.ga(&#39;toppmeny&#39;, &#39;sälja bostad&#39;, &#39;Utfälld meny&#39;);">
          Sälja bostad
          <i class="primary-nav__icon fa fa-angle-down"></i>
        </a>

        <div class="primary-nav-dropdown primary-nav-dropdown--sell">
          <ul class="primary-nav-dropdown__link-list">
            <li class="primary-nav-dropdown__link-list-item">
              <a class="primary-nav-dropdown__link qa-primary-nav-sell-property-market" onclick="hnt.ga(&#39;toppmeny&#39;, &#39;annonsera bostad&#39;, &#39;Annonsera bostad&#39;);" href="/annonsera-bostad">Annonsera på Hemnet</a>
              <span class="primary-nav-dropdown__highlight-text">Nyhet!</span>
            </li>
            <li class="primary-nav-dropdown__link-list-item">
              <a class="primary-nav-dropdown__link qa-primary-nav-sell-property-sell" onclick="hnt.ga(&#39;toppmeny&#39;, &#39;sälja bostad&#39;, &#39;Sälja bostad&#39;);" href="/salja-bostad">Guide till att sälja bostad</a>
            </li>
          </ul>
        </div>
      </li>
    <li class="primary-nav__list-item">
      <a href="/artiklar"
         class="primary-nav__link"
        onclick="hnt.ga(&#39;toppmeny&#39;, &#39;artiklar&#39;, this.href);">
        Artiklar
      </a>
    </li>
    <li class="primary-nav__list-item">
      <a href="/bygga-hus"
         class="primary-nav__link"
         onclick="hnt.ga(&#39;toppmeny&#39;, &#39;bygga hus&#39;, this.href);">
        Bygga hus
      </a>
    </li>
    <li class="primary-nav__list-item">
      <a href="/utland/"
         class="primary-nav__link"
         onclick="hnt.ga(&#39;toppmeny&#39;, &#39;utland&#39;, this.href);">
        Utland
      </a>
    </li>
    <li class="primary-nav__list-item primary-nav__list-item--mobile-only">
      <a
        class="primary-nav__link"
        href="/om/kontakt"
        onclick="hnt.ga(&#39;toppmeny&#39;, &#39;kontakta hemnet&#39;, this.href);">
        Kontakta Hemnet
      </a>
    </li>
  </ul>
    <ul class="secondary-nav secondary-nav--user-actions">
        <li class="secondary-nav__fixed">
          <a class="button button--secondary button--full-width" rel="nofollow" href="/mitt_hemnet/session/logga_in">Logga in</a>
        </li>
        <li class="secondary-nav__fixed">
          <a class="button button--primary button--full-width" rel="nofollow" href="/mitt_hemnet/anvandare/ny">Skapa konto</a>
        </li>
    </ul>
</nav>

  </div>

    <div class="site-header__user-nav">
      
<nav class="user-nav">
  <ul class="user-nav__list">
    <li class="user-link user-nav__list-item js-user-link" style="display: none;">
      <a class="user-nav__link" rel="nofollow" href="/mitt_hemnet/sparade_sokningar">Sparade sökningar</a>
    </li>
    <li class="user-link user-nav__list-item js-user-link" style="display: none;">
      <a class="user-nav__link" rel="nofollow" href="/mitt_hemnet/sparade_objekt">Sparade objekt</a>
    </li>
    <li class="user-link user-nav__list-item user-nav__list-item--settings js-user-link" style="display: none;">
      <a class="user-nav__link" rel="nofollow" href="/mitt_hemnet/sparade_sokningar">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path fill="#AAA" fill-rule="evenodd" d="M12 0c6.63 0 12 5.37 12 12 0 6.59-5.344 12-12 12-6.643 0-12-5.397-12-12C0 5.37 5.37 0 12 0zm8.29 18.094A10.34 10.34 0 0 0 22.286 12c0-5.665-4.62-10.286-10.286-10.286C6.335 1.714 1.714 6.334 1.714 12c0 2.277.75 4.38 1.996 6.094.482-2.398 1.647-4.38 4.098-4.38a5.98 5.98 0 0 0 8.384 0c2.45 0 3.616 1.982 4.098 4.38zM17.143 9.43A5.144 5.144 0 0 1 12 14.57a5.144 5.144 0 0 1-5.143-5.14A5.144 5.144 0 0 1 12 4.286a5.144 5.144 0 0 1 5.143 5.143z"/></svg>
        <span class="user-nav__name-label js-welcome-message" data-hj-masked data-user-id="">
          
        </span>
</a>    </li>
    <li class="anonymous-link user-nav__list-item js-anonymous-link" style="display: inline-block;">
      <a id="login-action" class="user-nav__link" rel="nofollow" href="/mitt_hemnet/session/logga_in">Logga in</a>
    </li>
    <li class="anonymous-link user-nav__list-item js-anonymous-link" style="display: inline-block;">
      <a class="user-nav__link" rel="nofollow" href="/mitt_hemnet/anvandare/ny">Skapa konto</a>
    </li>
  </ul>
</nav>

    </div>
</div>


    <div class="listing-container qa-listing-container">
    <div class="listing">
      



<div class="listing__panorama-banner">
  
  <div id="ad_unit_placeholder" class="listing__panorama-ad-placeholder">
  </div>
  <div id="ad-unit-0" class="ad-unit ad-unit--inline ">
  <script>
    googletag.cmd.push(function() { googletag.display("ad-unit-0"); });
  </script>
</div>


</div>

<div class="listing__property-info qa-property-info">
  
  <div class="property-info__container" style="">
    <div class="property-gallery">
      <div class="property-gallery__image-container">
  <div class="property-gallery__image ">
    <div class="property-carousel js-carousel-container qa-carousel" data-image-alt-text="1 rum Lägenhet på Jämtlandsgatan 152, 6 tr, Stockholm Vällingby Centrum" data-property-id="16663146">
        <div class="carousel__wrapper js-listing-gallery-initial-image">
          <img
            class="property-gallery__item"
            src=https://bilder.hemnet.se/images/itemgallery_cut/ca/16/ca162533055df0661ab81070ffe8d308.jpg
            alt=
          />
        </div>
        <div
          class="carousel__gallery-wrapper js-new-carousel"
          data-image-alt-text="1 rum Lägenhet på Jämtlandsgatan 152, 6 tr, Stockholm Vällingby Centrum"
          data-property-id="16663146"
        >
          <img
            class="property-gallery__item"
            src=https://bilder.hemnet.se/images/itemgallery_cut/ca/16/ca162533055df0661ab81070ffe8d308.jpg
            alt=
          />
        </div>
    </div>
    <div class="property-gallery__broker-logo">
        <img class="property-gallery__broker-logo--small" alt="Mäklarhuset Vällingby-Hässelby" src=https://bilder.hemnet.se/images/broker_logo_2/45/12/4512f7eb77c66e721ba9d52777e74d52.png />
        <img class="property-gallery__broker-logo--large" alt="Mäklarhuset Vällingby-Hässelby" src=https://bilder.hemnet.se/images/broker_logo/fe/1f/fe1f63d2c1860edbd928f0e8a49a955c.png />
    </div>
  </div>

</div>

<div
  class="js-actionbar-container"
  data-property-id="16663146"
  data-tracking-category="objektsida"
  data-is-saved="false"
  data-number-images="2"
  data-nfl="false"
></div>


    </div>

    <div class="property-info__top-labels">
    </div>

    <div class="property-info__content">
      <section class="property-info__data-container">
        <div class="property-info__primary-container">
          <div class="property-info__address-container">
            
<div class="property-address">
  <h1 class="property-address__street qa-property-heading">Jämtlandsgatan 152, 6 tr</h1>
  <div class="property-address__area-container">
    <span class="property-address__area">Vällingby Centrum, Stockholm</span>
      <a href="#karta" onclick="hnt.ga(&#39;objektsida&#39;, &#39;ankarlänk karta och gatuvy&#39;);">Se på karta</a>
  </div>
</div>

          </div>
          <div class="property-info__price-container">
            <p class="property-info__price qa-property-price">1 875 000 kr</p>
              <div class="property-info__price-change">
                <span class="property-info__original-price">1 995 000 kr</span>
                <span class="property-info__price-change-percentage">-6%</span>
              </div>
            <div class="property-info__bidding">
              <div class="property-info__bidding-label">
              </div>
            </div>
              <div
  class="property-info__watch-selling-price js-watch-selling-price"
  data-property="{&quot;id&quot;:&quot;16663146&quot;,&quot;street_address&quot;:&quot;Jämtlandsgatan 152, 6 tr&quot;}"
  data-broker-agency="{&quot;id&quot;:&quot;23247&quot;,&quot;name&quot;:&quot;Mäklarhuset Vällingby-Hässelby&quot;,&quot;websiteUrl&quot;:&quot;https://www.maklarhuset.se/vallingby&quot;,&quot;compactLogoUrl&quot;:&quot;https://bilder.hemnet.se/images/broker_logo_2/45/12/4512f7eb77c66e721ba9d52777e74d52.png&quot;}"
  data-broker="{&quot;name&quot;:&quot;Robert Blomster&quot;,&quot;email&quot;:&quot;robert.blomster@maklarhuset.se&quot;,&quot;phoneNumber&quot;:&quot;0732-313802&quot;,&quot;formattedPhoneNumber&quot;:&quot;+46732313802&quot;}"
  data-user="{&quot;loggedIn&quot;:false,&quot;email&quot;:null,&quot;phone&quot;:null,&quot;firstname&quot;:null,&quot;lastname&quot;:null}">
</div>

          </div>
        </div>
      </section>
    </div>
    <div class="property-info__attributes-and-description">
      

<div class="property-attributes">
  
<div class="property-attributes-table">
    <dl class="property-attributes-table__area">
        <div class="property-attributes-table__row">
          <dt class="property-attributes-table__label">Bostadstyp</dt>
          <dd class="property-attributes-table__value">Lägenhet</dd>
        </div>

        <div class="property-attributes-table__row">
          <dt class="property-attributes-table__label">Upplåtelseform</dt>
          <dd class="property-attributes-table__value">Bostadsrätt</dd>
        </div>

        <div class="property-attributes-table__row">
          <dt class="property-attributes-table__label">Antal rum</dt>
          <dd class="property-attributes-table__value">1 rum</dd>
        </div>

        <div class="property-attributes-table__row qa-living-area-attribute">
          <dt class="property-attributes-table__label">Boarea</dt>
          <dd class="property-attributes-table__value">33 m²</dd>
        </div>






        <div class="property-attributes-table__row">
          <dt class="property-attributes-table__label">Byggår</dt>
          <dd class="property-attributes-table__value">1953</dd>
        </div>


        <div class="property-attributes-table__row">
          <dt class="property-attributes-table__label">Förening</dt>
          <dd class="property-attributes-table__value">
            <div class="property-attributes-table__housing-cooperative-name">
              <span class="property-attributes-table__value">BRF Fornkullen</span>
            </div>
            <a href="#brf-info" onclick="">Om föreningen</a>
          </dd>
        </div>
    </dl>

  <dl class="property-attributes-table__area">
      <div class="property-attributes-table__row">
        <dt class="property-attributes-table__label">Avgift</dt>
        <dd class="property-attributes-table__value">1 723 kr/mån</dd>
      </div>

      <div class="property-attributes-table__row">
        <dt class="property-attributes-table__label">Driftkostnad</dt>
        <dd class="property-attributes-table__value">4 800 kr/år</dd>
      </div>



      <div class="property-attributes-table__row">
        <dt class="property-attributes-table__label">Pris/m²</dt>
        <dd class="property-attributes-table__value">56 818 kr/m²</dd>
      </div>

      <div class="property-attributes-table__row">
        <dt class="property-attributes-table__label"></dt>
        <dd class="property-attributes-table__value"><a href="#kalkyler" onclick="hnt.ga(&#39;bokostnad&#39;, &#39;visa kalkyler&#39;, &#39;ankarlänk under fakta&#39;);">Räkna på boendet</a></dd>
      </div>
  </dl>
</div>

</div>

<div class="property-description-container js-property-description-container qa-property-description-container">
  <div class="property-description-container__labels">
   </div>
  <div class="property-description js-property-description property-description--long">
    <p>Etta om 33 kvm i populära brf Fornkullen nära T-bana och Vällingby C.  Möjlig att göra om till en tvåa. Balkong i västerläge med fantastiks utsikt. Lägenheten målades om 2019.   Under 2018 har föreningen bytt fönster och balkongdörrar. Förening med mycket bra ekonomi, låg månadsavgift. Lugnt och centralt läge. Hiss och tvättstuga i huset. St. Stambytt helkaklat badrum. Utmärkt läge endast ca 200 meter till Vällingby C med butiker, restauranger, biograf, T-bana/buss och gym. Nära till simhall, Grimsta Naturreservat med motionsspår, badplatser, minigolf och fiske vid Mälaren. Välkommen på visning!</p>
  </div>
    <div class="property-description-expander-container">
      <button class="property-description-expander-button js-property-description-expander-button" data-target-element=".property-description">
         <i class="fa fa-chevron-down"></i> Visa hela beskrivningen
      </button>
    </div>
    <a class="property-description-broker-button qa-property-description-broker-button js-property-description-broker-button" target="_blank" rel="nofollow noopener" title="Alla bilder, planritning, dokument och längre beskrivning av objektet hos Mäklarhuset Vällingby-Hässelby" onclick="hnt.ga(&#39;objekt mäklarnavigation&#39;, &#39;utgående mäklarlänk 23247 Mäklarhuset Vällingby-Hässelby&#39;, &#39;objekt beskrivning&#39;);" href="https://www.maklarhuset.se/bostad/sverige/stockholm/vallingby/jamtlandsgatan/365523?referrer=hemnet">
        Läs mer hos mäklaren
      <i class="fa fa-external-link"></i>
</a></div>

    </div>
      <a class="broker-banner" title="Mäklarhuset Vällingby-Hässelby" target="_blank" rel="nofollow noopener" onclick="hnt.ga(&#39;objekt mäklarnavigation&#39;, &#39;utgående mäklarlänk 23247 Mäklarhuset Vällingby-Hässelby&#39;, &#39;objektbanner&#39;);" href="https://www.maklarhuset.se/vardering?utm_source=hemnet&amp;utm_medium=banner&amp;utm_campaign=vardering&amp;utm_content=bokakostnadsfrivardering">
    <img alt="Mäklarhuset Vällingby-Hässelby" src=https://bilder.hemnet.se/images/broker_banner/8e/00/8e007896a36e5f9d91bc8cc76595cfab.gif />
</a>
  </div>
  <div class="property-info__report-listing">
    <button
      data-target="#report-listing"
      class="report-listing__link js-report-listing__link"
      onclick="hnt.ga(&#39;objektsida&#39;, &#39;rapportera annonsfel öppna modal&#39;);"
      data-property-id="16663146"
      data-property-address="Jämtlandsgatan 152, 6 tr"
    >
      Rapportera fel i bostadsannonsen
    </button>
  </div>

</div>

<div class="listing__sidebar" id="visningar">
  <div class="listing__sidebar-container">
  <section class="listing__broker-info">
    

  <div class="broker-contact-card--sidebar">
  <h2 class="broker-section-heading qa-broker-section-heading ">
      Kontakta mäklaren
  </h2>

  <div class="broker-contact-card__content qa-broker">
  <div class="broker-contact-card__inner-content broker-contact-card__inner-content--sidebar">
    <div class="broker-contact-card__information">
      <p>
        <strong>
          Robert Blomster
        </strong>
      </p>
      <p>
          <a class="broker-link" target="_blank" rel="nofollow noopener" onclick="hnt.ga(&#39;objekt mäklarnavigation&#39;, &#39;utgående mäklarlänk 23247 Mäklarhuset Vällingby-Hässelby&#39;, &#39;mäklarinfo&#39;);" href="https://www.maklarhuset.se/vallingby">
            Mäklarhuset Vällingby-Hässelby <i class="fa fa-external-link external-link--desktop-only"></i>
</a>      </p>

      <div class="broker-contacts">
    <p class="broker-contact js-broker-contact qa-broker-contact">
      <span class="broker-contact__info-container">
        <span class="broker-contact__info">
          <span class="hide-element">Telefon:</span>
          <a onclick="hnt.ga(&#39;objektsida&#39;, &#39;telefon mäklare mobil&#39;);" class="broker-contact__link" href="tel:+46732313802">
            <i class="fa fa-phone"></i>0732-313802
</a>        </span>

        <button
          type="button"
          class="broker-contact__toggle-button gtm-tracking-broker-phone"
          onclick="hnt.ga(&#39;objektsida&#39;, &#39;visa telefon mäklare&#39;);"
          data-broker-id="5729"
        >
          Visa telefon
        </button>
      </span>

      <a onclick="hnt.ga(&#39;objektsida&#39;, &#39;telefon mäklare mobil&#39;);" class="broker-contact__button-link gtm-tracking-broker-phone" data-broker-id="5729" href="tel:+46732313802">
        Ring
</a>    </p>

    <p class="broker-contact js-broker-contact qa-broker-contact">
      <span class="broker-contact__info-container">
        <span class="broker-contact__info">
          <span class="hide-element">E-post:</span>
          <a class="broker-contact__link" onclick="hnt.ga(&#39;objektsida&#39;, &#39;mailadress mäklare&#39;);" data-broker-id="5729" href="/cdn-cgi/l/email-protection#bdcfd2dfd8cfc993dfd1d2d0cec9d8cffdd0dcd6d1dccfd5c8ced8c993ced882dfd2d9c480988dfc988dfc988dfcf5d8d0d3d8c990dcd3d3d2d3ced8d3988f8ddb98fe8e98ff8bcf988f8df798fe8e98fc89d0c9d1dcd3d9cedadcc9dcd3988f8d8c888f988ffe988f8d8b988f8dc9cf988efc988f8d988dfcd5c9c9cdce988efc988ffb988ffbcacaca93d5d8d0d3d8c993ced8988ffbdfd2cec9dcd9988ffbd1dcdad8d3d5d8c9908ccfc8d090cbdcd1d1d4d3dadfc490ded8d3c9cfc8d090cec9d2ded6d5d2d1d0ce90d6d2d0d0c8d390d7dcd0c9d1dcd3d9cedadcc9dcd3908c888f988ffe908b90c9cf908c8b8b8b8e8c898b9bdcd0cd86cec8dfd7d8dec980ebd4dc988f8df5d8d0d3d8c9988efc988f8dfcd3da98fe8e98fc88d8d3d9d8988f8df798fe8e98fc89d0c9d1dcd3d9cedadcc9dcd3988f8d8c888f988ffe988f8d8b988f8dc9cf">
            <i class="fa fa-envelope icon-color-grey"></i><span class="__cf_email__" data-cfemail="f98b969b9c8b8dd79b9596948a8d9c8bb994989295988b918c8a9c8dd78a9c">[email&#160;protected]</span>
</a>        </span>

        <button
          type="button"
          class="broker-contact__toggle-button gtm-tracking-broker-email"
          onclick="hnt.ga(&#39;objektsida&#39;, &#39;visa mailadress mäklare&#39;);"
          data-broker-id="5729"
        >
          Visa e-post
        </button>
      </span>

      <a class="broker-contact__button-link gtm-tracking-broker-email" onclick="hnt.ga(&#39;objektsida&#39;, &#39;mailadress mäklare&#39;);" data-broker-id="5729" href="/cdn-cgi/l/email-protection#b3c1dcd1d6c1c79dd1dfdcdec0c7d6c1f3ded2d8dfd2c1dbc6c0d6c79dc0d68cd1dcd7ca8e9683f29683f29683f2fbd6deddd6c79ed2dddddcddc0d6dd968183d596f08096f185c1968183f996f08096f287dec7dfd2ddd7c0d4d2c7d2dd9681838286819681f096818385968183c7c19680f29681839683f2dbc7c7c3c09680f29681f59681f5c4c4c49ddbd6deddd6c79dc0d69681f5d1dcc0c7d2d79681f5dfd2d4d6dddbd6c79e82c1c6de9ec5d2dfdfdaddd4d1ca9ed0d6ddc7c1c6de9ec0c7dcd0d8dbdcdfdec09ed8dcdedec6dd9ed9d2dec7dfd2ddd7c0d4d2c7d2dd9e8286819681f09e859ec7c19e828585858082878595d2dec388c0c6d1d9d6d0c78ee5dad2968183fbd6deddd6c79680f2968183f2ddd496f08096f286d6ddd7d6968183f996f08096f287dec7dfd2ddd7c0d4d2c7d2dd9681838286819681f096818385968183c7c1">
        Skicka e-post
</a>    </p>
</div>



      
<p class="">
  <a class="qa-broker-agency-listings-link" onclick="hnt.ga(&#39;objektsida&#39;, &#39;klick på Till salu från samma mäklarkontor&#39;, &#39;23247 Mäklarhuset Vällingby-Hässelby&#39;);" href="/maklare/objekt/23247">
    Till salu från
      samma kontor
</a></p>

    </div>
  </div>

</div>

</div>



  </section>
    <section class="listing__showings">
      

  <h2 class="listing-showings__header">Kom på visning</h2>
  <ul class="listing-showings__list qa-showings-list">
      <li class="listing-showings__showing qa-showings-item">

        <div class="listing-showings__calendar">
          <div>25</div>
          <div>mar</div>
        </div>

        <div class="listing-showings__showing-info">
          <span class="listing-showings__showing-time">
            Imorgon
          </span>
          <div class="listing-showings__showing-description">
            Välkommen på visning – för tider,mer info och anmälan gå till Mäklarhusets sida
          </div>
        </div>
        <a title="Ladda ner visningstiden till din personliga digitala kalender (.ics-format)." class="listing-showings__showing-link listing-showings__showing-link--small qa-showing-button" onclick="hnt.ga(&#39;objektsida&#39;, &#39;lägg till i kalender&#39;, &#39;objekt: 16663146&#39;);" href="https://www.hemnet.se/bostad/lagenhet-1rum-vallingby-centrum-stockholms-kommun-jamtlandsgatan-152,-6-tr-16663146/visning/26737596.ics">
            <i class="listing-showings__icon fa fa-calendar-plus-o" aria-hidden="true" aria-label="Lägg till i kalender"></i>
</a>      </li>
      <li class="listing-showings__showing qa-showings-item">

        <div class="listing-showings__calendar">
          <div>29</div>
          <div>mar</div>
        </div>

        <div class="listing-showings__showing-info">
          <span class="listing-showings__showing-time">
            Sön 29 mar
          </span>
          <div class="listing-showings__showing-description">
            Välkommen på visning – för tider,mer info och anmälan gå till Mäklarhusets sida
          </div>
        </div>
        <a title="Ladda ner visningstiden till din personliga digitala kalender (.ics-format)." class="listing-showings__showing-link listing-showings__showing-link--small qa-showing-button" onclick="hnt.ga(&#39;objektsida&#39;, &#39;lägg till i kalender&#39;, &#39;objekt: 16663146&#39;);" href="https://www.hemnet.se/bostad/lagenhet-1rum-vallingby-centrum-stockholms-kommun-jamtlandsgatan-152,-6-tr-16663146/visning/26737735.ics">
            <i class="listing-showings__icon fa fa-calendar-plus-o" aria-hidden="true" aria-label="Lägg till i kalender"></i>
</a>      </li>
  </ul>
  <span class="listing-showings__information">För bokad eller privat visning sker anmälan hos mäklaren</span>

    </section>
    <section class="listing__sidebar-ads">
      
    </section>
    <section class="listing__calculator-anchors">
      <div class="calculator-anchors">
  <div class="calculator-anchors__header-container">
    <h2 class="calculator-anchors__header">Vad blir din boendekostnad?</h2>
  </div>
  <ul class="calculator-anchors__list">
    <li onclick="hnt.ga(&#39;bokostnad&#39;, &#39;visa kalkyler&#39;, &#39;högerkolumn&#39;);" class="calculator-anchors__list-item">
      <a class="js-calculator-anchor-link" data-tab-id="living_cost" href="#kalkyler">Se bankernas kalkyler</a>
    </li>
    <li onclick="hnt.ga(&#39;försäkringskostnad&#39;, &#39;visa kalkyler&#39;, &#39;högerkolumn&#39;);" class="calculator-anchors__list-item">
      <a class="js-calculator-anchor-link" data-tab-id="insurance" href="#kalkyler">Försäkringskostnad</a>
    </li>
    <li onclick="hnt.ga(&#39;låna till annat&#39;, &#39;visa kalkyler&#39;, &#39;högerkolumn&#39;);" class="calculator-anchors__list-item">
      <a class="js-calculator-anchor-link" data-tab-id="other_loan" href="#kalkyler">Låna till annat</a>
    </li>
    <li onclick="hnt.ga(&#39;bredbandskalkyl&#39;, &#39;visa kalkyler&#39;, &#39;högerkolumn&#39;);" class="calculator-anchors__list-item">
      <a class="js-calculator-anchor-link" data-tab-id="broadband_cost" href="#kalkyler">Se erbjudanden för bredband/TV</a>
    </li>
  </ul>
</div>

    </section>
</div>

</div>

<div class="listing__showing-top-callout">
  


      <a class="showing-top-callout showing-top-callout--showing" href="#visningar" onclick="hnt.ga(&#39;objektsida&#39;, &#39;callout&#39;, &#39;Nästa visning Imorgon &#39;)">
        <div class="listing-showings__calendar">
          <div>25</div>
          <div>mar</div>
        </div>
        <div class="showing-top-callout__container">
            <p class="showing-top-callout__title">Kom på visning</p>
            <p class="showing-top-callout__date">Imorgon </p>
        </div>
        <p class="showing-top-callout__more">Mer info</p>
      </a>

</div>

<div class="listing__related-info-container">

  <section class="listing__calculators">


  <div id="kalkyler">
    
<div class="section-heading">
  <div class="section-heading__strike">
    <div class="section-heading__strike--left"></div>
  </div>
  <div class="section-heading__content">
      <h2 class="section-heading__text">Räkna på ditt nya boende</h2>
    
  </div>
  <div class="section-heading__strike">
    <div class="section-heading__strike--right"></div>
  </div>
</div>


    <div class="js-calculators-skeleton">
      <div class="calculators-skeleton__content">
      </div>
    </div>

    <div class="js-calculators" data-property-id="16663146">
    </div>
  </div>
</section>

  <section class="listing__housing-cooperative" id="brf-info">
    



  
<div class="section-heading">
  <div class="section-heading__strike">
    <div class="section-heading__strike--left"></div>
  </div>
  <div class="section-heading__content">
      <h2 class="section-heading__text">Om BRF Fornkullen</h2>
    
  </div>
  <div class="section-heading__strike">
    <div class="section-heading__strike--right"></div>
  </div>
</div>

  <div class="js-housing-cooperative qa-housing-cooperative" data-property-id="16663146"></div>
  <div class="js-housing-cooperative-skeleton">
    <div class="housing-cooperative-skeleton__content"></div>
    <div class="housing-cooperative-skeleton__disclaimer"></div>
  </div>

  </section>
  <section class="listing__maps" id="karta">
    

  <div id="link-anchor-map">
    
<div class="section-heading">
  <div class="section-heading__strike">
    <div class="section-heading__strike--left"></div>
  </div>
  <div class="section-heading__content">
      <h2 class="section-heading__text">Karta &amp; restider</h2>
    
  </div>
  <div class="section-heading__strike">
    <div class="section-heading__strike--right"></div>
  </div>
</div>

      <div id="property-water-information" class="property-water">
        <div class="property-water-info">
          <div class="property-water-info__image">
            <svg xmlns="http://www.w3.org/2000/svg" class="property-water-info__svg" width="128" height="50" viewBox="0 0 128 50">
             <g fill-rule="evenodd">
               <path d="M58.79 7.446h3.9v-5.53h-3.9v5.53zM37 18.576h32.764l8.38-12.94H64.18V.428H57.3v5.205H44.706L37 18.575zm49.69.05l-7.616-12.35-.026-.04L71 18.646l15.69-.02zM39.08 35.318h48.067c8.433 2.95 11.1 7.46 12.23 9.405H21.974c.808-1.413 2.47-3.89 4.76-5.13 3.32-1.797 4.552-2.458 12.344-4.275zm-1.758-1.11c-6.736 1.617-8.062 2.328-11.3 4.08-3.744 2.025-5.868 6.677-5.956 6.874-.104.23-.084.495.054.707.137.21.374.336.626.336h80.54c.307 0 .58-.184.692-.468.114-.283.044-.605-.177-.814-.33-.314-3.632-7.305-13.59-10.707V19.963l-18.05.092v13.777H58.15v-7.577h-5.22v7.577H38.813v-13.81H37.32v14.185zm66.828 9.705h16.19V43h-16.19v.913zm-9.19 6h16.19V49H94.96v.913zm-86.468 0h39.27V49H8.492v.913z"/>
               <path d="M0 39.913h24.27V39H0v.913zm94 0h33.27V39H94v.913z"/>
             </g>
            </svg>
          </div>
          <div class="property-water-info__details">
              <p class="property-water-info__details-text">
                Ungefär 1,1 km till vatten (Råcksta träsk)
              </p>
              <p class="property-water-info__details-text">
                Ungefär 7,5 km till havet
              </p>
          </div>
        </div>
      </div>

    <div class="js-map-skeleton listing-page-map  listing-page-map--large"></div>
    <div class="js-react-listing-map" id="map" data-property-id=16663146
      data-show-schools=true data-large-map=true></div>


  </div>

  </section>
  <section class="listing__sellers-page" id="tjanster">
    
  <div class="js-sellers-page-upsell-skeleton sellers-page-upsell-skeleton qa-sellers-page-upsell-skeleton">
    <div class="sellers-page-upsell-skeleton__title"></div>
    <div class="sellers-page-upsell-skeleton__subtitle"></div>
    <div class="sellers-page-upsell-skeleton__visits"></div>
    <div class="sellers-page-upsell-skeleton__cta"></div>
    <div class="sellers-page-upsell-skeleton__visits-text"></div>
    <div class="sellers-page-upsell-skeleton__products"></div>
  </div>

  <div class="js-sellers-page-upsell qa-sellers-page-upsell" id="sellers-page-upsell" data-property-id="16663146" />

  </section>
  <section class="listing__price-trend">
    


  <div class="js-price-trend-skeleton">
    
<div class="section-heading">
  <div class="section-heading__strike">
    <div class="section-heading__strike--left"></div>
  </div>
  <div class="section-heading__content">
      <h2 class="section-heading__text">Prisutveckling i Stockholms kommun</h2>
    
  </div>
  <div class="section-heading__strike">
    <div class="section-heading__strike--right"></div>
  </div>
</div>

    <div class="price-trend-skeleton__content">
      <div class="price-trend-skeleton__period-select"></div>
      <div class="price-trend-skeleton__monthly-change"></div>
      <div class="price-trend-skeleton__graph"></div>
    </div>
  </div>

  <div class="js-price-trend" id="price-trend" data-property-id="16663146">
  </div>


  </section>
  <section class="listing__area-information">
    <div class="area-information__container">
  <div class="area-information__icon">
    <svg width="339" height="298" viewBox="0 0 339 298" xmlns="http://www.w3.org/2000/svg"><title>illustration_step9</title><g fill="none" fill-rule="evenodd"><path d="M316.935 149c0 82.29-66.71 149-149 149s-149-66.71-149-149 66.71-149 149-149 149 66.71 149 149" fill="#F6F6F6"/><path d="M262.178 198.502s-5.384.216-11.036 8.372c-5.65 8.158-10.863 12.59-10.863 12.59s-17.096-14.46-27.85-27.082c0 0-2.417-.282-5.36 2.247-2.944 2.528-4.476.95-4.817 2.886-.34 1.936 12.64 32.635 38.497 44.69 0 0 13.756-3.052 21.428-9.932" fill="#EBF6EA"/><path d="M262.178 198.502s-5.384.216-11.036 8.372c-5.65 8.158-10.863 12.59-10.863 12.59s-17.096-14.46-27.85-27.082c0 0-2.417-.282-5.36 2.247-2.944 2.528-4.476.95-4.817 2.886-.34 1.936 12.64 32.635 38.497 44.69 0 0 13.756-3.052 21.428-9.932" stroke="#000" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/><path fill="#FFF" d="M285.983 192.65l50.696 83.2h-73.95v-78.218"/><path stroke="#000" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" d="M285.983 192.65l50.696 83.2h-73.95v-78.218"/><path fill="#39AA35" d="M262.73 192.65v4.982h26.338l-3.5-5.753"/><path stroke="#000" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" d="M262.73 192.65v4.982h26.338l-3.5-5.753"/><path d="M285.568 190.643c-11.486-5.75-5.264-23.75-5.264-23.75 5.25-21.5-7.982-27.58-7.982-27.58-22.233-12.956-21.72-21.31-21.72-21.31s-7.98 2.933-7.98 18.89c0 16.133-10.498 13.82-9.253 21.175 4.152.505 11.566 0 11.566 0s-1.424 4.686-.158 8.957c0 0 1.463 2.392 5.733.178 0 0 3.19-.712 1.84 2.195-.963 2.073-4.626 2.55-3.618 4.685 1.01 2.136 4.567 3.38 4.33 5.754-.237 2.372-2.61 12.04 4.687 12.04h27.818" fill="#FFF"/><path d="M285.568 190.643c-11.486-5.75-5.264-23.75-5.264-23.75 5.25-21.5-7.982-27.58-7.982-27.58-22.233-12.956-21.72-21.31-21.72-21.31s-7.98 2.933-7.98 18.89c0 16.133-10.498 13.82-9.253 21.175 4.152.505 11.566 0 11.566 0s-1.424 4.686-.158 8.957c0 0 1.463 2.392 5.733.178 0 0 3.19-.712 1.84 2.195-.963 2.073-4.626 2.55-3.618 4.685 1.01 2.136 4.567 3.38 4.33 5.754-.237 2.372-2.61 12.04 4.687 12.04h27.818" stroke="#000" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/><path d="M76.766 198.502s5.383.216 11.034 8.372c5.653 8.158 10.584 14.96 10.584 14.96s10.683-6.957 18.428-15.027c5.322-5.545 8.105-12.55 9.702-14.425 0 0 2.415-.282 5.358 2.247 2.944 2.528 4.476.95 4.818 2.886.34 1.936-12.64 32.635-38.497 44.69 0 0-14.757-3.724-22.43-10.604" fill="#F5FBF5"/><path d="M76.766 198.502s5.383.216 11.034 8.372c5.653 8.158 10.584 14.96 10.584 14.96s10.683-6.957 18.428-15.027c5.322-5.545 8.105-12.55 9.702-14.425 0 0 2.415-.282 5.358 2.247 2.944 2.528 4.476.95 4.818 2.886.34 1.936-12.64 32.635-38.497 44.69 0 0-14.757-3.724-22.43-10.604" stroke="#000" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/><path fill="#F5FBF5" d="M75.29 198.107l.475 77.744H1.5l47.928-78.455"/><path stroke="#000" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" d="M75.29 198.107l.475 77.744H1.5l47.928-78.455"/><path fill="#D6EDD8" d="M75.764 191.86v5.535H49.427l3.075-5.935"/><path stroke="#000" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" d="M75.764 191.86v5.535H49.427l3.075-5.935"/><path d="M87.05 116.004s11.244 4.053 11.244 20.01c0 16.134 8.13 14.164 6.886 21.52-4.153.505-11.568 0-11.568 0s1.424 4.686.16 8.957c0 0-1.464 2.393-5.735.178 0 0-3.188-.712-1.84 2.194.964 2.075 4.628 2.55 3.62 4.687-1.01 2.135-4.568 3.38-4.33 5.753.237 2.372 2.61 12.04-4.686 12.04-7.296 0-29.052.12-29.052.12" fill="#FFF"/><path d="M87.05 116.004s11.244 4.053 11.244 20.01c0 16.134 8.13 14.164 6.886 21.52-4.153.505-11.568 0-11.568 0s1.424 4.686.16 8.957c0 0-1.464 2.393-5.735.178 0 0-3.188-.712-1.84 2.194.964 2.075 4.628 2.55 3.62 4.687-1.01 2.135-4.568 3.38-4.33 5.753.237 2.372 2.61 12.04-4.686 12.04-7.296 0-29.052.12-29.052.12M198.906 182.447s-3.954-4.745-6.09-1.897c-2.135 2.846 1.582 5.06 4.034 6.01" stroke="#000" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/><path d="M256.296 145.715c0 1.434-1.162 2.595-2.596 2.595-1.432 0-2.594-1.16-2.594-2.595 0-1.434 1.162-2.595 2.594-2.595 1.434 0 2.596 1.16 2.596 2.595" fill="#000"/><path d="M250.63 117.673s-1.32-16.537 13.695-16.537 16.794 14.224 16.794 14.224 30.703-.284 30.703 26.818c0 27.1-7.3 43.058-25.022 49.82 0 0-8.784-1.506-8.784-13.536s10.015-30.044-7.26-40.137c-17.278-10.093-20.125-13.78-20.125-20.652" fill="#C4E4C3"/><path d="M250.63 117.673s-1.32-16.537 13.695-16.537 16.794 14.224 16.794 14.224 30.703-.284 30.703 26.818c0 27.1-7.3 43.058-25.022 49.82 0 0-8.784-1.506-8.784-13.536s10.015-30.044-7.26-40.137c-17.278-10.093-20.125-13.78-20.125-20.652zM211.007 191.463s0-5.457-7.04-9.017M202.246 197.08s-.282-2.577-1.52-3.085M127.935 191.463s0-5.457 7.04-9.017M139.16 183.1s4.83-5.398 6.966-2.55c2.136 2.846-1.58 5.06-4.033 6.01M136.696 197.08s.283-2.577 1.52-3.085" stroke="#000" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/><path d="M82.254 147.957c0 1.433 1.16 2.595 2.595 2.595 1.432 0 2.593-1.162 2.593-2.595 0-1.434-1.16-2.596-2.594-2.596-1.435 0-2.596 1.163-2.596 2.597" fill="#000"/><path d="M78.53 161.287c0 2.95-2.394 5.342-5.344 5.342s-5.342-2.393-5.342-5.343 2.392-5.343 5.342-5.343c2.95 0 5.343 2.393 5.343 5.343" fill="#39AA35"/><path d="M78.53 161.287c0 2.95-2.394 5.342-5.344 5.342s-5.342-2.393-5.342-5.343 2.392-5.343 5.342-5.343c2.95 0 5.343 2.393 5.343 5.343z" stroke="#000" stroke-width="3"/><path d="M45.495 98.448c0 8.823-7.152 15.976-15.976 15.976-8.824 0-15.977-7.153-15.977-15.976 0-8.824 7.153-15.977 15.976-15.977s15.975 7.154 15.975 15.978" fill="#D6EDD8"/><path d="M45.495 98.448c0 8.823-7.152 15.976-15.976 15.976-8.824 0-15.977-7.153-15.977-15.976 0-8.824 7.153-15.977 15.976-15.977s15.975 7.154 15.975 15.978z" stroke="#000" stroke-width="3"/><path d="M200.726 177.312c0 15.303-14.546 27.708-32.487 27.708-17.942 0-32.487-12.405-32.487-27.708 0-15.304 14.545-27.71 32.487-27.71 17.94 0 32.486 12.406 32.486 27.71z" stroke="#39AA35" stroke-width="6" stroke-linecap="round" stroke-linejoin="round"/><path fill="#39AA35" d="M195.543 32.256V26.88H171.22V18.85h-6.204v130.75h6.203V65.43h24.324v-5.38h-7.71V32.257"/><g><path d="M195.543 186.956s.907 2.53 0 4.508c0 0 4.11 3.32 6.703 1.74 2.592-1.582 1.96-4.27.536-4.825-1.424-.554-4.588-2.77-6.485-1.424" fill="#F5FBF5"/><path d="M195.543 186.956s.907 2.53 0 4.508c0 0 4.11 3.32 6.703 1.74 2.592-1.582 1.96-4.27.536-4.825-1.424-.554-4.588-2.77-6.485-1.424" stroke="#000" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/></g><g><path d="M143.4 186.956s-.908 2.53 0 4.508c0 0-4.112 3.32-6.704 1.74-2.592-1.582-1.96-4.27-.535-4.825 1.425-.554 4.588-2.77 6.486-1.424" fill="#FFF"/><path d="M143.4 186.956s-.908 2.53 0 4.508c0 0-4.112 3.32-6.704 1.74-2.592-1.582-1.96-4.27-.535-4.825 1.425-.554 4.588-2.77 6.486-1.424" stroke="#000" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/></g><path d="M87.05 116.004s2.76 20.925-25.767 34.29c0 0-9.695 6.06-8.815 38.558" stroke="#F00" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/><path d="M87.05 116.004s-32.488-23.49-52.816 0c-20.33 23.49-11.943 75.455 18.27 75.455 0 0-1.168-34.203 9.87-40.975 11.035-6.77 28.615-13.444 24.675-34.48" fill="#D6EDD8"/><path d="M87.05 116.004s-32.488-23.49-52.816 0c-20.33 23.49-11.943 75.455 18.27 75.455 0 0-1.168-34.203 9.87-40.975 11.035-6.77 28.615-13.444 24.675-34.48z" stroke="#000" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/></g></svg>
  </div>
  <div class="area-information__text">
    Är du nyfiken på att veta mer om området?
  </div>
  <div class="area-information__link">
    <a class="area-information__button qa-area-information-link" rel="nofollow" href="/omrade/stockholms-kommun">Omdömen, skolor och mer prisinformation</a>
  </div>
</div>

  </section>
  <section class="listing__similar-listings">
      
  <div class="similar-listings qa-similar-listings">
    <div class="similar-listings__heading">
      
<div class="section-heading">
  <div class="section-heading__strike">
    <div class="section-heading__strike--left"></div>
  </div>
  <div class="section-heading__content">
      <h2 class="section-heading__text">Liknande bostäder till salu</h2>
    
  </div>
  <div class="section-heading__strike">
    <div class="section-heading__strike--right"></div>
  </div>
</div>

    </div>
    <ul class="similar-listings__list">
          <li class="similar-listings__list-item">
            <a class="similar-listing" onclick="hnt.ga(&#39;objektsida&#39;, &#39;klick liknande objekt&#39;, 1);" href="https://www.hemnet.se/bostad/lagenhet-1rum-vallingby-centrum-stockholms-kommun-jamtlandsgatan-150-16686595">
  <div class="similar-listing__image">
      <img class="js-lazy-load" data-src=https://bilder.hemnet.se/images/itemgallery_M/75/7b/757b7d59892359c2e3cfb2e064ba4c2f.jpg />
    <div class="similar-listing__state-labels-container">
        <div class="listing-card__state-label-container">
           <span class="listing-card__label listing-card__state-label highlighted_showing">
            
<span class="state-label-icon--highlighted_showing">
  <i class="fa calendar-white-icon"></i>
</span>

            Sön 29 mar kl 13:00
          </span>
        </div>
    </div>
  </div>
  <div class="similar-listing-content">
    <div class="similar-listing__location-container">
      <p class="similar-listing__address"><svg width="14" height="16" viewBox="0 0 14 16" xmlns="http://www.w3.org/2000/svg"><title>Lägenhet</title><desc><span class="svg-icon__fallback-text">Lägenhet</span></desc><path class="svg-icon__shape" d="M0 1.333v13.334C0 15.403.597 16 1.333 16h4v-5.333H8V16h4c.737 0 1.333-.597 1.333-1.333V1.333C13.333.597 12.737 0 12 0H1.333C.597 0 0 .597 0 1.333" fill="#C1569D" fill-rule="evenodd"/></svg>
Jämtlandsgatan 150</p>
      <p>
        <span>Vällingby Centrum, Stockholm</span>
      </p>
    </div>
    <p>
      <span>1 695 000 kr</span>
    </p>
    <p>
      <span>33 m²</span>
      <span>1 rum</span>
    </p>
    <p>
      <span></span>
    </p>
    <div class="similar-listing-card__labels-container">
          <span class="similar-listing-card__label similar-listing-card__label--PRODUCT">
            Plus
          </span>
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Balkong
          </span>
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Hiss
          </span>
    </div>
  </div>
</a>

          </li>
          <li class="similar-listings__list-item">
            <a class="similar-listing" onclick="hnt.ga(&#39;objektsida&#39;, &#39;klick liknande objekt&#39;, 2);" href="https://www.hemnet.se/bostad/lagenhet-1rum-bromma-blackeberg-stockholms-kommun-bjornsonsgatan-142-16699635">
  <div class="similar-listing__image">
      <img class="js-lazy-load" data-src=https://bilder.hemnet.se/images/itemgallery_M/22/dc/22dc954c2177f3177c82dab3257c173d.jpg />
    <div class="similar-listing__state-labels-container">
        <div class="listing-card__state-label-container">
           <span class="listing-card__label listing-card__state-label highlighted_showing">
            
<span class="state-label-icon--highlighted_showing">
  <i class="fa calendar-white-icon"></i>
</span>

            Sön 29 mar kl 15:15
          </span>
        </div>
    </div>
  </div>
  <div class="similar-listing-content">
    <div class="similar-listing__location-container">
      <p class="similar-listing__address"><svg width="14" height="16" viewBox="0 0 14 16" xmlns="http://www.w3.org/2000/svg"><title>Lägenhet</title><desc><span class="svg-icon__fallback-text">Lägenhet</span></desc><path class="svg-icon__shape" d="M0 1.333v13.334C0 15.403.597 16 1.333 16h4v-5.333H8V16h4c.737 0 1.333-.597 1.333-1.333V1.333C13.333.597 12.737 0 12 0H1.333C.597 0 0 .597 0 1.333" fill="#C1569D" fill-rule="evenodd"/></svg>
Björnsonsgatan 142</p>
      <p>
        <span>Bromma Blackeberg, Stockholm</span>
      </p>
    </div>
    <p>
      <span>1 790 000 kr</span>
    </p>
    <p>
      <span>31 m²</span>
      <span>1 rum</span>
    </p>
    <p>
      <span></span>
    </p>
    <div class="similar-listing-card__labels-container">
          <span class="similar-listing-card__label similar-listing-card__label--PRODUCT">
            Plus
          </span>
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Balkong
          </span>
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Hiss
          </span>
    </div>
  </div>
</a>

          </li>
          <li class="similar-listings__list-item">
            <a class="similar-listing" onclick="hnt.ga(&#39;objektsida&#39;, &#39;klick liknande objekt&#39;, 3);" href="https://www.hemnet.se/bostad/lagenhet-1,5rum-vallingby-stockholms-kommun-jamtlandsgatan-152,-8tr-16670242">
  <div class="similar-listing__image">
      <img class="js-lazy-load" data-src=https://bilder.hemnet.se/images/itemgallery_M/b0/a7/b0a76c7699b4df15023dd3c68b945218.jpg />
    <div class="similar-listing__state-labels-container">
    </div>
  </div>
  <div class="similar-listing-content">
    <div class="similar-listing__location-container">
      <p class="similar-listing__address"><svg width="14" height="16" viewBox="0 0 14 16" xmlns="http://www.w3.org/2000/svg"><title>Lägenhet</title><desc><span class="svg-icon__fallback-text">Lägenhet</span></desc><path class="svg-icon__shape" d="M0 1.333v13.334C0 15.403.597 16 1.333 16h4v-5.333H8V16h4c.737 0 1.333-.597 1.333-1.333V1.333C13.333.597 12.737 0 12 0H1.333C.597 0 0 .597 0 1.333" fill="#C1569D" fill-rule="evenodd"/></svg>
Jämtlandsgatan 152, 8tr</p>
      <p>
        <span>Vällingby, Stockholm</span>
      </p>
    </div>
    <p>
      <span>1 695 000 kr</span>
    </p>
    <p>
      <span>33 m²</span>
      <span>1,5 rum</span>
    </p>
    <p>
      <span></span>
    </p>
    <div class="similar-listing-card__labels-container">
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Balkong
          </span>
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Hiss
          </span>
    </div>
  </div>
</a>

          </li>
          <li class="similar-listings__list-item">
            <a class="similar-listing" onclick="hnt.ga(&#39;objektsida&#39;, &#39;klick liknande objekt&#39;, 4);" href="https://www.hemnet.se/bostad/lagenhet-1rum-vallingby-centrum-stockholms-kommun-jamtlandsgatan-72-16521068">
  <div class="similar-listing__image">
      <img class="js-lazy-load" data-src=https://bilder.hemnet.se/images/itemgallery_M/68/18/68187b96949f085705029d23d771c71b.jpg />
    <div class="similar-listing__state-labels-container">
    </div>
  </div>
  <div class="similar-listing-content">
    <div class="similar-listing__location-container">
      <p class="similar-listing__address"><svg width="14" height="16" viewBox="0 0 14 16" xmlns="http://www.w3.org/2000/svg"><title>Lägenhet</title><desc><span class="svg-icon__fallback-text">Lägenhet</span></desc><path class="svg-icon__shape" d="M0 1.333v13.334C0 15.403.597 16 1.333 16h4v-5.333H8V16h4c.737 0 1.333-.597 1.333-1.333V1.333C13.333.597 12.737 0 12 0H1.333C.597 0 0 .597 0 1.333" fill="#C1569D" fill-rule="evenodd"/></svg>
Jämtlandsgatan 72</p>
      <p>
        <span>Vällingby Centrum, Stockholm</span>
      </p>
    </div>
    <p>
      <span>1 495 000 kr</span>
    </p>
    <p>
      <span>23 m²</span>
      <span>1 rum</span>
    </p>
    <p>
      <span></span>
    </p>
    <div class="similar-listing-card__labels-container">
    </div>
  </div>
</a>

          </li>
          <li class="similar-listings__list-item">
            <a class="similar-listing" onclick="hnt.ga(&#39;objektsida&#39;, &#39;klick liknande objekt&#39;, 5);" href="https://www.hemnet.se/bostad/lagenhet-1rum-vallingby-centrum-stockholms-kommun-morsilsgatan-8-16669201">
  <div class="similar-listing__image">
      <img class="js-lazy-load" data-src=https://bilder.hemnet.se/images/itemgallery_M/ac/91/ac91970e5d6991703cc066e27ce97b4f.jpg />
    <div class="similar-listing__state-labels-container">
    </div>
  </div>
  <div class="similar-listing-content">
    <div class="similar-listing__location-container">
      <p class="similar-listing__address"><svg width="14" height="16" viewBox="0 0 14 16" xmlns="http://www.w3.org/2000/svg"><title>Lägenhet</title><desc><span class="svg-icon__fallback-text">Lägenhet</span></desc><path class="svg-icon__shape" d="M0 1.333v13.334C0 15.403.597 16 1.333 16h4v-5.333H8V16h4c.737 0 1.333-.597 1.333-1.333V1.333C13.333.597 12.737 0 12 0H1.333C.597 0 0 .597 0 1.333" fill="#C1569D" fill-rule="evenodd"/></svg>
Mörsilsgatan 8</p>
      <p>
        <span>Vällingby Centrum, Stockholm</span>
      </p>
    </div>
    <p>
      <span>1 695 000 kr</span>
    </p>
    <p>
      <span>41 m²</span>
      <span>1 rum</span>
    </p>
    <p>
      <span></span>
    </p>
    <div class="similar-listing-card__labels-container">
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Balkong
          </span>
    </div>
  </div>
</a>

          </li>
          <li class="similar-listings__list-item">
            <a class="similar-listing" onclick="hnt.ga(&#39;objektsida&#39;, &#39;klick liknande objekt&#39;, 6);" href="https://www.hemnet.se/bostad/lagenhet-1rum-vallingby-parkstad-racksta-stockholms-kommun-jamtlandsgatan-107,-van-8-11084-16673975">
  <div class="similar-listing__image">
      <img class="js-lazy-load" data-src=https://bilder.hemnet.se/images/itemgallery_M/46/1a/461ae362f45ee79e2b365dde62dbb1eb.jpg />
    <div class="similar-listing__state-labels-container">
        <div class="listing-card__state-label-container">
           <span class="listing-card__label listing-card__state-label ongoing_bidding">
            
<span class="state-label-icon--ongoing_bidding">
  <i class="fa fa-legal"></i>
</span>

            Budgivning pågår
          </span>
        </div>
    </div>
  </div>
  <div class="similar-listing-content">
    <div class="similar-listing__location-container">
      <p class="similar-listing__address"><svg width="14" height="16" viewBox="0 0 14 16" xmlns="http://www.w3.org/2000/svg"><title>Lägenhet</title><desc><span class="svg-icon__fallback-text">Lägenhet</span></desc><path class="svg-icon__shape" d="M0 1.333v13.334C0 15.403.597 16 1.333 16h4v-5.333H8V16h4c.737 0 1.333-.597 1.333-1.333V1.333C13.333.597 12.737 0 12 0H1.333C.597 0 0 .597 0 1.333" fill="#C1569D" fill-rule="evenodd"/></svg>
Jämtlandsgatan 107, Vån 8 (11084)</p>
      <p>
        <span>Vällingby Parkstad / Råcksta, Stockholm</span>
      </p>
    </div>
    <p>
      <span>1 895 000 kr</span>
    </p>
    <p>
      <span>36 m²</span>
      <span>1 rum</span>
    </p>
    <p>
      <span></span>
    </p>
    <div class="similar-listing-card__labels-container">
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Balkong
          </span>
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Hiss
          </span>
    </div>
  </div>
</a>

          </li>
          <li class="similar-listings__list-item">
            <a class="similar-listing" onclick="hnt.ga(&#39;objektsida&#39;, &#39;klick liknande objekt&#39;, 7);" href="https://www.hemnet.se/bostad/lagenhet-1rum-vallingby-racksta-stockholms-kommun-jamtlandsgatan-97-16695680">
  <div class="similar-listing__image">
      <img class="js-lazy-load" data-src=https://bilder.hemnet.se/images/itemgallery_M/93/39/9339e1730269eab083bc3972aab78ca2.jpg />
    <div class="similar-listing__state-labels-container">
    </div>
  </div>
  <div class="similar-listing-content">
    <div class="similar-listing__location-container">
      <p class="similar-listing__address"><svg width="14" height="16" viewBox="0 0 14 16" xmlns="http://www.w3.org/2000/svg"><title>Lägenhet</title><desc><span class="svg-icon__fallback-text">Lägenhet</span></desc><path class="svg-icon__shape" d="M0 1.333v13.334C0 15.403.597 16 1.333 16h4v-5.333H8V16h4c.737 0 1.333-.597 1.333-1.333V1.333C13.333.597 12.737 0 12 0H1.333C.597 0 0 .597 0 1.333" fill="#C1569D" fill-rule="evenodd"/></svg>
Jämtlandsgatan 97</p>
      <p>
        <span>Vällingby Råcksta, Stockholm</span>
      </p>
    </div>
    <p>
      <span>1 895 000 kr</span>
    </p>
    <p>
      <span>36 m²</span>
      <span>1 rum</span>
    </p>
    <p>
      <span></span>
    </p>
    <div class="similar-listing-card__labels-container">
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Uteplats
          </span>
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Hiss
          </span>
    </div>
  </div>
</a>

          </li>
          <li class="similar-listings__list-item">
            <a class="similar-listing" onclick="hnt.ga(&#39;objektsida&#39;, &#39;klick liknande objekt&#39;, 8);" href="https://www.hemnet.se/bostad/lagenhet-1,5rum-vallingby-racksta-stockholms-kommun-ostersundsgatan-2,-4-tr-16422957">
  <div class="similar-listing__image">
      <img class="js-lazy-load" data-src=https://bilder.hemnet.se/images/itemgallery_M/38/e9/38e91a1e9caaeb73722d6fa4832005c1.jpg />
    <div class="similar-listing__state-labels-container">
    </div>
  </div>
  <div class="similar-listing-content">
    <div class="similar-listing__location-container">
      <p class="similar-listing__address"><svg width="14" height="16" viewBox="0 0 14 16" xmlns="http://www.w3.org/2000/svg"><title>Lägenhet</title><desc><span class="svg-icon__fallback-text">Lägenhet</span></desc><path class="svg-icon__shape" d="M0 1.333v13.334C0 15.403.597 16 1.333 16h4v-5.333H8V16h4c.737 0 1.333-.597 1.333-1.333V1.333C13.333.597 12.737 0 12 0H1.333C.597 0 0 .597 0 1.333" fill="#C1569D" fill-rule="evenodd"/></svg>
Östersundsgatan 2, 4 tr</p>
      <p>
        <span>Vällingby Råcksta, Stockholm</span>
      </p>
    </div>
    <p>
      <span>1 995 000 kr</span>
    </p>
    <p>
      <span>42 m²</span>
      <span>1,5 rum</span>
    </p>
    <p>
      <span></span>
    </p>
    <div class="similar-listing-card__labels-container">
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Balkong
          </span>
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Hiss
          </span>
    </div>
  </div>
</a>

          </li>
          <li class="similar-listings__list-item">
            <a class="similar-listing" onclick="hnt.ga(&#39;objektsida&#39;, &#39;klick liknande objekt&#39;, 9);" href="https://www.hemnet.se/bostad/lagenhet-1rum-norra-angby-stockholms-kommun-bjorketorpsvagen-13-16612270">
  <div class="similar-listing__image">
      <img class="js-lazy-load" data-src=https://bilder.hemnet.se/images/itemgallery_M/7f/04/7f04da70db303ab0c28acc8d07ae6bb6.jpg />
    <div class="similar-listing__state-labels-container">
        <div class="listing-card__state-label-container">
           <span class="listing-card__label listing-card__state-label ongoing_bidding">
            
<span class="state-label-icon--ongoing_bidding">
  <i class="fa fa-legal"></i>
</span>

            Budgivning pågår
          </span>
        </div>
    </div>
  </div>
  <div class="similar-listing-content">
    <div class="similar-listing__location-container">
      <p class="similar-listing__address"><svg width="14" height="16" viewBox="0 0 14 16" xmlns="http://www.w3.org/2000/svg"><title>Lägenhet</title><desc><span class="svg-icon__fallback-text">Lägenhet</span></desc><path class="svg-icon__shape" d="M0 1.333v13.334C0 15.403.597 16 1.333 16h4v-5.333H8V16h4c.737 0 1.333-.597 1.333-1.333V1.333C13.333.597 12.737 0 12 0H1.333C.597 0 0 .597 0 1.333" fill="#C1569D" fill-rule="evenodd"/></svg>
Björketorpsvägen 13</p>
      <p>
        <span>Norra Ängby, Stockholm</span>
      </p>
    </div>
    <p>
      <span>1 900 000 kr</span>
    </p>
    <p>
      <span>28 m²</span>
      <span>1 rum</span>
    </p>
    <p>
      <span></span>
    </p>
    <div class="similar-listing-card__labels-container">
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Hiss
          </span>
    </div>
  </div>
</a>

          </li>
          <li class="similar-listings__list-item">
            <a class="similar-listing" onclick="hnt.ga(&#39;objektsida&#39;, &#39;klick liknande objekt&#39;, 10);" href="https://www.hemnet.se/bostad/lagenhet-1rum-bromma-blackeberg-stockholms-kommun-bjornsonsgatan-138-16670975">
  <div class="similar-listing__image">
      <img class="js-lazy-load" data-src=https://bilder.hemnet.se/images/itemgallery_M/44/bc/44bcb0b86b301559dd81fd118228b08f.jpg />
    <div class="similar-listing__state-labels-container">
        <div class="listing-card__state-label-container">
           <span class="listing-card__label listing-card__state-label upcoming_live_stream_showing">
            
<span class="state-label-icon--upcoming_live_stream_showing">
  <i class="fa fa-play-circle"></i>
</span>

            Onlinevisning imorgon kl 17:00
          </span>
        </div>
    </div>
  </div>
  <div class="similar-listing-content">
    <div class="similar-listing__location-container">
      <p class="similar-listing__address"><svg width="14" height="16" viewBox="0 0 14 16" xmlns="http://www.w3.org/2000/svg"><title>Lägenhet</title><desc><span class="svg-icon__fallback-text">Lägenhet</span></desc><path class="svg-icon__shape" d="M0 1.333v13.334C0 15.403.597 16 1.333 16h4v-5.333H8V16h4c.737 0 1.333-.597 1.333-1.333V1.333C13.333.597 12.737 0 12 0H1.333C.597 0 0 .597 0 1.333" fill="#C1569D" fill-rule="evenodd"/></svg>
Björnsonsgatan 138</p>
      <p>
        <span>Bromma / Blackeberg, Stockholm</span>
      </p>
    </div>
    <p>
      <span>1 725 000 kr</span>
    </p>
    <p>
      <span>31 m²</span>
      <span>1 rum</span>
    </p>
    <p>
      <span></span>
    </p>
    <div class="similar-listing-card__labels-container">
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Balkong
          </span>
          <span class="similar-listing-card__label similar-listing-card__label--FEATURE">
            Hiss
          </span>
    </div>
  </div>
</a>

          </li>
    </ul>
    <div class='similar-listings-search-link'><a class="similar-listings-search-link__button" onclick="hnt.ga(&#39;objektsida&#39;, &#39;länk till fler liknande objekt&#39;, this.getAttribute(&#39;href&#39;));" href="https://www.hemnet.se/bostader?item_types%5B%5D=bostadsratt&amp;location_ids%5B%5D=473464&amp;rooms_max=1.5&amp;rooms_min=0.5">Visa lägenheter till salu i Vällingby</a></div>
  </div>

  </section>
  <div class="listing__breadcrumbs">
    
<div class="listing-breadcrumbs-box">
  <p class="listing-breadcrumbs-box__description">
    Kanske finns det andra lägenheter  du har missat?
  </p>
  <ul class="breadcrumbs">
      <li class="breadcrumbs__item">
        <i class="breadcrumbs__icon fa fa-angle-right"></i>
        <a class="breadcrumbs__link" onclick="hnt.ga(&#39;objektsida&#39;, &#39;brödsmulor&#39;, &#39;Län&#39;);" href="/bostader?housing_form_groups%5B%5D=apartments&amp;location_ids%5B%5D=17744">Stockholms län</a>
        (4172)
      </li>
      <li class="breadcrumbs__item">
        <i class="breadcrumbs__icon fa fa-angle-right"></i>
        <a class="breadcrumbs__link" onclick="hnt.ga(&#39;objektsida&#39;, &#39;brödsmulor&#39;, &#39;Kommun&#39;);" href="/bostader?housing_form_groups%5B%5D=apartments&amp;location_ids%5B%5D=18031">Stockholms kommun</a>
        (1918)
      </li>
      <li class="breadcrumbs__item">
        <i class="breadcrumbs__icon fa fa-angle-right"></i>
        <a class="breadcrumbs__link" onclick="hnt.ga(&#39;objektsida&#39;, &#39;brödsmulor&#39;, &#39;Område&#39;);" href="/bostader?housing_form_groups%5B%5D=apartments&amp;location_ids%5B%5D=925953">Hässelby - Vällingby</a>
        (65)
      </li>
  </ul>
</div>

  </div>
  <section class="listing__articles">


<div class="section-heading">
  <div class="section-heading__strike">
    <div class="section-heading__strike--left"></div>
  </div>
  <div class="section-heading__content">
      <h2 class="section-heading__text">Artiklar från Hemnet</h2>
    
  </div>
  <div class="section-heading__strike">
    <div class="section-heading__strike--right"></div>
  </div>
</div>



<div class="js-articles-skeleton">
  <div class="articles-skeleton__container">
    <div class="articles-skeleton__article"></div>
    <div class="articles-skeleton__article"></div>
    <div class="articles-skeleton__article"></div>
  </div>
</div>

<div class="js-listing-articles" data-property-id="16663146">
</div>
</section>
</div>

    </div>
  </div>

  <div class="footer-bar qa-footer-bar-sellers-guide js-footer-bar">
  <div class="footer-bar__sellers-guide-illustration"></div>
  <a href="/salja-bostad" class="footer-bar__cta-link" onclick="hnt.ga(&#39;Sälja bostad&#39;, &#39;övriga interna länkar&#39;, &#39;puff sidfot&#39;);">
    <div class="footer-bar__heading">
      Går du i säljtankar?
    </div>
    <div class="footer-bar__subheading">
      Hemnet guidar dig till en lyckad försäljning &raquo;
    </div>
  </a>
</div>

  <footer>
  <div class="footer-sitemap js-footer qa-site-footer">
    <div class="footer-sitemap__wrapper">
      <div class="footer-sitemap__column footer-sitemap__column--contact">
        <h2 class="footer-sitemap__heading">Kontakta oss</h2>
          <ul class="footer-sitemap__list">
            <li class="footer-sitemap__list-item">
              <a href="/om/kontakt" onclick="hnt.ga('sidfot', 'kontakta oss');">Kontakta Hemnet</a>
            </li>
            <li class="footer-sitemap__list-item">
              <a href="/om" onclick="hnt.ga('sidfot', 'om');">Om Hemnet</a>
            </li>
            <li class="footer-sitemap__list-item">
              <a href="https://press.hemnet.se" onclick="hnt.ga('sidfot', 'press');">Press</a>
            </li>
            <li class="footer-sitemap__list-item">
              <a href="/statistik" onclick="hnt.ga('sidfot', 'bostadsstatistik');">Bostadsstatistik</a>
            </li>
            <li class="footer-sitemap__list-item">
              <a href="https://reklam.hemnet.se" onclick="hnt.ga('sidfot', 'kop reklamplats');">Köp reklamplats</a>
            </li>
            <li class="footer-sitemap__list-item">
              <a href="https://jobba.hemnet.se/" onclick="hnt.ga('sidfot', 'jobba hos oss');">Jobba hos oss</a>
            </li>
            <li class="footer-sitemap__list-item">
              <a href="https://press.hemnet.se/trender" onclick="hnt.ga('sidfot', 'om bostadsmarknaden');">Om bostadsmarknaden</a>
            </li>
          </ul>
      </div>

      <div class="footer-sitemap__column footer-sitemap__column--apps">
        <h2 class="footer-sitemap__heading">Hemnet i mobilen</h2>

        <div class="footer-sitemap__container"> 
          <i class="left fa fa-mobile fa-5x footer__icon-mobile" style="margin: -4px 0 0 9px;">&nbsp;</i>
          <ul class="footer-sitemap__list">
            <li class="footer-sitemap__list-item">
              <a href="https://itunes.apple.com/app/apple-store/id525017304?pt=272428&ct=hemnet-se-footer&mt=8" onclick="hnt.ga('sidfot', 'ladda ner ios');" rel="nofollow">
                Ladda ner iPhone-/iPad-appen<br>
              </a>
            </li>
            <li class="footer-sitemap__list-item">
              <a href="https://play.google.com/store/apps/details?id=se.hemnet.android"
                onclick="hnt.ga('sidfot', 'ladda ner android');" rel="nofollow">
                Ladda ner Android-appen
              </a>
            </li>
          </ul>
        </div>
      </div>

      <div class="footer-sitemap__column footer-sitemap__column--social">
        <h2 class="footer-sitemap__heading">Följ Hemnet</h2>
        <ul class="footer-sitemap__list">
          <li class="clear-children">
            <a href="/nyhetsbrev" onclick="hnt.ga('sidfot', 'nyhetsbrev');" rel="nofollow">
              <span class="fa-stack fa-lg">
                <i class="fa fa-circle fa-stack-2x footer__icon-envelope">&nbsp;</i>
                <i class="fa fa-envelope fa-stack-1x fa-inverse">&nbsp;</i>
              </span>
              <span class="social-icon-label">Läs vårt nyhetsbrev</span>
            </a>
          </li>
          <li class="clear-children">
            <a href="https://www.facebook.com/hemnet.se" rel="nofollow" onclick="hnt.ga('sidfot', 'facebook');">
              <span class="fa-stack fa-lg">
                <i class="fa fa-circle fa-stack-2x footer__icon-facebook">&nbsp;</i>
                <i class="fa fa-facebook fa-stack-1x fa-inverse">&nbsp;</i>
              </span>
              <span class="social-icon-label">Gilla oss på Facebook</span>
            </a>
          </li>
          <li class="clear-children">
            <a href="https://www.twitter.com/hemnet" rel="nofollow" onclick="hnt.ga('sidfot', 'twitter');">
              <span class="fa-stack fa-lg">
                <i class="fa fa-circle fa-stack-2x footer__icon-twitter">&nbsp;</i>
                <i class="fa fa-twitter fa-stack-1x fa-inverse">&nbsp;</i>
              </span>
              <span class="social-icon-label">Följ oss på Twitter</span>
            </a>
          </li>
          <li class="clear-children">
            <a href="https://www.instagram.com/hemnet" rel="nofollow" onclick="hnt.ga('sidfot', 'instagram');">
              <span class="fa-stack fa-lg">
                <i class="fa fa-circle fa-stack-2x footer__icon-instagram">&nbsp;</i>
                <i class="fa fa-camera fa-stack-1x fa-inverse">&nbsp;</i>
              </span>
              <span class="social-icon-label">Följ oss på Instagram</span>
            </a>
          </li>
          <li class="clear-children">
            <a href="https://www.pinterest.se/hemnet" rel="nofollow" onclick="hnt.ga('sidfot', 'pinterest');">
              <span class="fa-stack fa-lg">
                <i class="fa fa-pinterest fa-stack-2x">&nbsp;</i>
              </span>
              <span class="social-icon-label">Hitta inspiration på Pinterest</span>
            </a>
          </li>
        </ul>
      </div>
    </div>
  </div>
  <div class="footer-smallprint__wrapper">
  <div class="footer-smallprint">
    <ul class="footer-smallprint__element footer-smallprint__link-list">
      <li><a href="/sitemap">Alla bostäder</a></li>
      <li><a href="/om/cookies" rel="nofollow">Så använder Hemnet cookies</a></li>
      <li><a href="/om/integritet" rel="nofollow">Integritet</a></li>
    </ul>

    <div class="footer-smallprint__element footer-smallprint__publishing-info">
      <p>
        <span class="footer-smallprint__text-block">Ansvarig utgivare: Jessica Sjöberg</span>
        <span class="footer-smallprint__text-block">© 2020 Hemnet AB</span>
      </p>
    </div>
  </div>
</div>

</footer>


    <!-- Snowplow -->
<script data-cfasync="false" src="/cdn-cgi/scripts/5c5dd728/cloudflare-static/email-decode.min.js"></script><script>
(function(p,l,o,w,i,n,g){if(!p[i]){p.GlobalSnowplowNamespace=p.GlobalSnowplowNamespace||[];
  p.GlobalSnowplowNamespace.push(i);p[i]=function(){(p[i].q=p[i].q||[]).push(arguments)
  };p[i].q=p[i].q||[];n=l.createElement(o);g=l.getElementsByTagName(o)[0];n.async=1;
  n.src=w;g.parentNode.insertBefore(n,g)}}(window,document,'script','https://assets.hemnet.se/assets/javascripts/vendor/snowplow-2.9.0.js', 'spTrack'));

window.spTrack('newTracker', 'webTracker', 'tracking.hemnet.se', {
  appId: 'hemnet-ng',
  platform: 'web',
  cookieDomain: 'tracking.hemnet.se',
  userFingerprint: true,
  forceSecureTracker: true,
  cookieLifetime: 60 * 60 * 24 * 31 * 24, // Two years
  stateStorageStrategy: 'localStorage',
  contexts: {
    webPage: true,
    gaCookies: true,
  }
});
</script>
<!-- end snowplow -->


<div class="js-google-tag-manager-id" data-tag-id="GTM-MMHTZ7"></div>
<!-- Google Tag Manager -->
<script>
  var userStatus;
  if (Hemnet.server.isUserLoggedIn()) {
    if (Hemnet.server.userHasVerifiedEmail()) {
      userStatus = 'registered verified user';
    } else {
      userStatus = 'registered unverified user';
    }
  } else {
    userStatus = 'guest';
  }

  dataLayer = [{"page":{"type":"standard"}},{"property":{"id":16663146,"broker_firm":"Mäklarhuset Vällingby-Hässelby","broker_agency_id":23247,"location":"Stockholm","locations":{"country":"Sverige","county":"Stockholms län","municipality":"Stockholms kommun","postal_city":"Vällingby","district":"Hässelby - Vällingby, Kista / Hässelby / Vällingby / Spånga, Stockholm, Vällingby","street":"Jämtlandsgatan","city":"Stockholm tätort"},"images_count":22,"new_production":false,"offers_selling_price":true,"status":"for_sale","housing_form":"Lägenhet","tenure":"Bostadsrätt","street_address":"Jämtlandsgatan 152, 6 tr","rooms":1.0,"living_area":33.0,"driftkostnad":4800,"price_per_m2":56818,"price":1875000,"has_price_change":true,"borattavgift":1723,"upcoming_open_houses":true,"bidding":false,"long_description":true,"live_streams":false,"live_stream_today":false,"has_floorplan":true,"publication_date":"2020-03-09","similar_items_length":10,"construction_year":"1953","housing_cooperative":"BRF Fornkullen","callout":"Imorgon ","amenities":["elevator"],"listing_package_type":"basic","experiment_features":[]}}];
  dataLayer.push({
    user: {
      status: userStatus,
      userId: Hemnet.server.userId(),
      shouldVerifyEmail: Hemnet.server.userShouldVerifyEmail()
    },
    isInternalTraffic: Hemnet.server.data.isInternalTraffic || false
  });
  dataLayer.push({experiments: "[{\"n\":\"kpis\",\"v\":\"A\"},{\"n\":\"sdh\",\"v\":\"A\"},{\"n\":\"nfl\",\"v\":\"A\"},{\"n\":\"lvc\",\"v\":\"A\"},{\"n\":\"nrm\",\"v\":\"A\"}]"});
</script>
<noscript>
  <iframe src="https://www.googletagmanager.com/ns.html?id=GTM-MMHTZ7"
    height="0" width="0" style="display:none;visibility:hidden"></iframe>
</noscript>

<script src="https://assets.hemnet.se/assets/packs/google-tag-manager.5fd4383df05e1a9c5c53.js" async="async"></script>

<!-- End Google Tag Manager -->

    <script src="https://assets.hemnet.se/assets/packs/polyfill.e77ccb4bd683b85c265d.js"></script>
    <script src="https://assets.hemnet.se/assets/packs/vendor.a768aa9a297896a8b3e0.js"></script>
      <script src="https://assets.hemnet.se/assets/packs/listing-page.83364e354b56e66a48b8.js" async="async"></script>
  

<script>
  window.loopaData = {"ProductId":"16663146","Price":"1875000","HousingForm":"APARTMENT","Tenure":"TENANT_OWNERSHIP","Rooms":"1.0","LivingArea":"33.0","Country":"Sverige","CountryId":"17691","Region":[],"RegionId":[],"Municipality":["Stockholms","Stockholms"],"MunicipalityId":["18031","18031"],"City":["Stockholm","Stockholm"],"CityId":["898734","898734"],"PostalCity":["Vällingby","Vällingby"],"PostalCityId":["18158","18158"],"District":["Vällingby","Kista / Hässelby / Vällingby / Spånga","Hässelby - Vällingby","Stockholm","Vällingby","Kista / Hässelby / Vällingby / Spånga","Hässelby - Vällingby","Stockholm"],"DistrictId":["473464","898747","925953","940128","473464","898747","925953","940128"],"Street":"Jämtlandsgatan","StreetId":"501997"};

  function isLocalStorageEnabled() {
    try {
      var testKey = '_is_localstorage_enabled';
      localStorage.setItem(testKey, testKey);
      localStorage.removeItem(testKey);
      return true;
    } catch (e) {
      return false;
    };
  };

  if (isLocalStorageEnabled()) {
    var signifiValue = localStorage.getItem("_signify_consent");
    if (signifiValue && signifiValue.includes("YES")) {
      var signifiUrl = "https://loopaautomate.azureedge.net/Hemnet.se/publisherLoopaAutomateInclude.js";

      if (signifiUrl.length > 0) {
        var script = document.createElement('script');
        script.src = signifiUrl;
        script.async = "async";
        document.body.appendChild(script);
      };
    };
  };
</script>


    <script src="https://assets.hemnet.se/assets/packs/listing-carousel.5a27c5c1ea99a171271a.js" async="async"></script>
  <script src="https://assets.hemnet.se/assets/packs/listing-description.3215cfa547691017c5fb.js" async="async"></script>
  <script src="https://assets.hemnet.se/assets/packs/broker-contact.5b704b1aebb47e55ac31.js"></script>
    <script src="https://assets.hemnet.se/assets/packs/listing-calculators.2bb09545c1d7f8540c2f.js" async="async"></script>
    <script src="https://assets.hemnet.se/assets/packs/housing-cooperative.e0b557079cab55279974.js" async="async"></script>
      <script src="https://assets.hemnet.se/assets/packs/listing-page-map.c1d542f5c942c8748188.js"></script>
  <script src="https://assets.hemnet.se/assets/packs/charts.d90edf33843d1d8c5b3c.js" async="async"></script>
  <script src="https://assets.hemnet.se/assets/packs/price-trend.95f63f59a1dd82bc10fa.js" async="async"></script>
  <script src="https://assets.hemnet.se/assets/packs/listing-articles.040d6a842fa52e2f740d.js" async="async"></script>
  <script src="https://assets.hemnet.se/assets/packs/top-navigation.5209660129459c517ac7.js" async="async"></script>

  </body>

</html>`

func Test_parseJson(t *testing.T) {
	tests := []struct {
		name string
		key  string
		want string
	}{
		{name: "postal_city", key: "postal_city", want: "Vällingby"},
		{name: "municipality", key: "municipality", want: "Stockholms kommun"},
		{name: "housing_form", key: "housing_form", want: "Lägenhet"},
		{name: "rooms", key: "rooms", want: "1.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseJson(tt.key, []byte(document)); got != tt.want {
				t.Errorf("parseJson() = %v, want %v", got, tt.want)
			}
		})
	}
}
