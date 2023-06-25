<script lang="ts">
  // @ts-ignore
  import { page } from "$app/stores";
  let error: Error = $page.error;

  import { onMount } from "svelte";

  onMount(() => {
    console.error("Something went wrong: ", error);
  });
</script>

<h1>Oops!</h1>
<h2>{error.message}</h2>
<p>
  This page probably hasnt been fully implemented yet. We appriciate your
  patience.
</p>
