import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
    title: "GetHooky",
    description: "CLI tool for managing your git hooks",
    base: "/GetHooky",
    head: [
        [
            "link",
            {
                rel: "icon",
                href: "./assets/getHooky.png"
            }
        ]
    ],
    themeConfig: {
        // https://vitepress.dev/reference/default-theme-config
        sidebar: [
            { text: "Introduction", link: "/" },
            { text: "Get Started", "link": "/get-started" },
            { text: "How To", "link": "/how-to" },
            { text: "Changelog", "link": "/changelog" }
        ],

        socialLinks: [
            { icon: 'github', link: 'https://github.com/ezpieco/gethooky' }
        ]
    }
})
