import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Tana.moe API",
  description: "A simple manga tracker back-end built with PocketBase",
  outDir: "dist",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    logo: "/logo.png",
    siteTitle: false,
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Guide', link: '/guide/getting-started' }
    ],

    sidebar: [
      {
        text: 'Introduction',
        items: [
          { text: 'Getting Started', link: '/guide/getting-started' },
        ]
      },
      {
        text: 'Reference',
        items: [
          { 
            text: 'Collections',
            items: [
              { text: 'titles', link: '/reference/collections/titles' }, 
            ]
          },
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/tanamoe' }
    ],

    search: {
      provider: 'local'
    }
  }
})
