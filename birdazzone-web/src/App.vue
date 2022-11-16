<script setup lang="ts">
import NavBar from "./components/NavBar.vue"
import Footer from "./components/Footer.vue"
</script>

<template>
  <div class="flex h-full">
    <NavBar />
    <RouterView/>
  </div>
  <Footer/>
</template>
