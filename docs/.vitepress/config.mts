import { defineConfig } from "vitepress";

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
      { text: "Home", link: "/" },
      { text: "Guide", link: "/guide/getting-started" },
    ],

    sidebar: [
      {
        text: "Introduction",
        items: [{ text: "Getting Started", link: "/getting-started" }],
      },
      {
        text: "Guides",
        items: [
          {
            text: "Metadata",
            link: "/guide/metadata",
          }
        ]
      },
      {
        text: "Collections",
        items: [
          {
            text: "titles",
            collapsed: true,
            link: "/collections/titles/index",
            items: [
              {
                text: "releases",
                link: "/collections/titles/releases",
              },
              {
                text: "publications",
                link: "/collections/titles/publications",
              },
              {
                text: "books",
                link: "/collections/titles/books",
              },
            ],
          },
        ],
      },
    ],

    socialLinks: [{ icon: "github", link: "https://github.com/tanamoe" }],

    search: {
      provider: "local",
    },
  },
});
