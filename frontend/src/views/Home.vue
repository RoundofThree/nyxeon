<template>
    <main class="w-full">
      <div
        class="relative pt-16 pb-32 flex content-center items-center justify-center"
        style="min-height: 50vh;"
      >
      <!-- Black background -->
        <div
          class="absolute top-0 w-full h-full bg-center bg-cover"
        >
          <span
            id="blackOverlay"
            class="w-full h-full absolute opacity-75 bg-black"
          ></span>
        </div>

        <div class="container relative mx-auto">
          <div class="items-center flex flex-wrap">
            <div class="w-full lg:w-6/12 px-4 ml-auto mr-auto text-center">
              <div class="pr-12">
                <h1 class="text-white font-semibold text-5xl">
                  Record your daily achievements
                </h1>
                <p class="mt-4 text-lg text-gray-300">
                  TBD: Explain better the purpose of this application. 
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <section class="pt-20 pb-48">
        <div class="container mx-auto px-4">
          <div class="flex flex-wrap justify-center text-center mb-24">
            <div class="w-full lg:w-6/12 px-4">
              <h2 class="text-4xl font-semibold">Sign in or sign up with: </h2>
              <button class="mt-10" v-on:click="githubLogin()">
                <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24"><path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/></svg>
              </button>
            </div>
          </div>
        </div>
      </section>

    </main>
</template>

<script>
// call Go backend api to return a URL  
export default {
  methods: {
    githubLogin() {
      window.location.href = "https://github.com/login/oauth/authorize?client_id=86fcb5a15a814ff63552&redirect_uri=http://localhost:8080/oauth/github/callback"
    }
  }
}
</script>
