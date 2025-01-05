import { defineConfig } from "astro/config";
import starlight from "@astrojs/starlight";

// https://astro.build/config
export default defineConfig({
  integrations: [
    starlight({
      title: "Blom",
      description:
        "Blom is a programming language built in Go, offering compilation with QBE, direct interpretation, and transpilation to multiple languages.",
      logo: {
        src: "./src/assets/logo.png",
      },
      favicon: "/favicon.png",
      social: {
        github: "https://github.com/xhyrom/blom",
      },
      lastUpdated: true,
      editLink: {
        baseUrl: "https://github.com/xhyrom/blom/tree/main/www",
      },
      customCss: ["./src/styles/index.css"],
      sidebar: [
        {
          label: "Guides",
          items: [{ label: "Example Guide", slug: "guides/example" }],
        },
        {
          label: "Reference",
          autogenerate: { directory: "reference" },
        },
      ],
    }),
  ],
});
