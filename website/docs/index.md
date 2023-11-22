---
title: Ablegram
titleTemplate: Search Ableton projects

layout: home

hero:
  name: Ablegram
  text: Search through your Ableton project files.
  # tagline: Search effortlessly through your Ableton project files.
  image:
    src: ./assets/screenshots/overview-v1.4.0.png
    alt: Overview of the Ablegram application
  actions:
    - theme: brand
      text: Get Started
      link: /introduction
    - theme: alt
      text: Download
      link: https://github.com/adrianrudnik/ablegram/releases/latest
    - theme: alt
      text: Try live demo
      link: https://demo.ablegram.app/

features:
  - title: Cross-platform
    details: Built to work across Windows, MacOS and Linux platforms.
  - title: Completely offline
    details: No Internet connection required, no subscription plans to worry about. 
  - title: Comprehensive indexing
    details: Full text index of your Ableton files - covering their internal structure.
  - title: Content tagging
    details: All findings are tagged. Search your files using text components or tags. 

---

<SchemaOrg>
{
    "@context": "https://schema.org",
    "@type": "SoftwareApplication",
    "name": "Ablegram",
    "operatingSystem": ["Windows 10", "Windows 11", "MacOS", "Linux"],
    "applicationCategory": "MultimediaApplication",
    "offers": {
        "@type": "Offer",
        "price": "0"
    }
}
</SchemaOrg>
