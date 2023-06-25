<script lang="ts">
  import "./+layout.css";

  // @ts-ignore
  import { page } from "$app/stores";
  let error: Error = $page.error;
</script>

<nav>
  <a href="/" aria-current={$page.url.pathname === "/"}> Home </a>

  <a href="/about" aria-current={$page.url.pathname.startsWith("/about")}>
    About
  </a>
</nav>

<slot />
