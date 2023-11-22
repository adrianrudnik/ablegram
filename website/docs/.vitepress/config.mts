import {defineConfig} from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Ablegram",
  description: "Search effortlessly through your Ableton project files.",

  head: [
    ['link', {rel: "shortcut icon", type: "image/x-icon", sizes: "16x16 32x32 48x48", href: "/favicon.ico"}],
    ['link', {rel: "icon", type: "image/png", sizes: "192x192", href: "/favicon-192.png"}],
  ],

  // Block shared snippets from becoming pages
  srcExclude: [
    './search/parser/shared'
  ],

  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config

    editLink: {
      pattern: 'https://github.com/adrianrudnik/ablegram/tree/main/website/docs/:path'
    },

    footer: {
      message: 'Website realized with <a href="https://vitepress.dev/" target="_blank">VitePress</a>',
    },

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
        collapsed: false,
        items: [
          {
            text: 'Examples',
            link: '/search/examples/',
          },
          {
            text: 'Parsers',
            link: '/search/parser/',
            items: [
              {text: 'Ableton Live Set', link: '/search/parser/ableton-live-set'},
              {text: 'Ableton Info-Text', link: '/search/parser/ableton-info-text'},
            ]
          },
          {
            text: 'Mappings',
            link: '/search/mapping-type/',
            collapsed: true,
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
            collapsed: true,
            items: [
              {text: 'Ableton .ALS', link: '/search/internals/ableton-als-file'},
            ]
          },
          {
            text: 'APIs',
            collapsed: true,
            items: [
              {text: 'Tags API', link: '/search/api/tags'},
            ]
          }
        ],
      },
      {
        text: 'Push 3',
        link: '/push3/',
        items: [
          {text: 'Web API', link: '/push3/web-api'},
        ]
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
