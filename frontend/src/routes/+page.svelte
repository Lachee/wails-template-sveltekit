<script lang="ts">
  import logo from "$lib/assets/images/logo-universal.png";
  import { Greeting } from "$wails/app/App";

  let nameText: string = "";
  let resultText: string = "👇 Enter your name below 👇";

  async function UpdateWelcome() {
    resultText = await Greeting(nameText);
    console.log("EEE", nameText, resultText);
  }
</script>

<main>
  <img alt="Wails logo" id="logo" src={logo} />
  <div class="result" id="result">{resultText}</div>
  <input
    type="text"
    on:change={() => UpdateWelcome()}
    bind:value={nameText}
    placeholder="Your Name..."
  />
</main>

<style>
  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }
</style>
