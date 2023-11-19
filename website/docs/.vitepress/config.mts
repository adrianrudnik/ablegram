import {defineConfig} from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Ablegram",
  titleTemplate: ':title',
  description: "Search effortlessly through your Ableton project files.",

  // Block shared snippets from becoming pages
  srcExclude: [
    './search/parser/shared'
  ],

  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config

    lastUpdated: true,

    nav: [
      {text: 'Home', link: '/'},
      {text: 'Privacy', link: '/legal/privacy-policy'},
      {text: 'Imprint', link: '/legal/imprint'},
    ],

    sidebar: [
      {
        text: 'Guide',
        items: [
          {text: 'Introduction', link: '/introduction'},
          {text: 'Installation', link: '/installation'},
          {text: 'Live Demo', link: '/live-demo'},
          {text: 'Service arguments', link: '/service-arguments'},
        ]
      },
      {
        text: 'Search',
        collapsed: true,
        items: [
          {
            text: 'Parsers',
            link: '/search/parser/',
            items: [
              {text: 'Ableton Live Set', link: '/search/parser/ableton-live-set'},
            ]
          },
          {
            text: 'Mappings',
            link: '/search/mapping-type/',
            items: [
              {text: 'Exact mapping', link: '/search/mapping-type/exact'},
              {text: 'Fulltext mapping', link: '/search/mapping-type/fulltext'},
              {text: 'Numerical mapping', link: '/search/mapping-type/numerical'},
              {text: 'Boolean mapping', link: '/search/mapping-type/boolean'},
            ]
          },
          {
            text: 'Internals',
            link: '/search/internals/',
            items: [
              {text: 'Ableton .ALS', link: '/search/internals/ableton-als-file'},
            ]
          }
        ],
      }
    ],

    search: {
      provider: 'local'
    },

    outline: {
      level: 'deep',
    },

    socialLinks: [
      {icon: 'github', link: 'https://github.com/adrianrudnik/ablegram'}
    ],
  },
  sitemap: {
    // https://vitepress.dev/guide/sitemap-generation
    hostname: 'https://www.ablegram.app',
    transformItems: (items) => {
      items = items.filter((item) =>
        !item.url.startsWith('_') &&
        !item.url.includes('/_') &&
        !item.url.startsWith('legal/') &&
        item.lastmod
      )

      return items
    }
  },
})
