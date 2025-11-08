<script setup lang="ts">
import { computed } from "vue";
import { useSubprocesses } from "../../queries/subprocess";
import Subprocess from "./Subprocess.vue";

const { data, asyncStatus } = useSubprocesses();

const sortedSubprocesses = computed(() => {
  if (data.value) {
    return [...data.value].sort((a, b) => a.pid - b.pid);
  } else {
    return undefined;
  }
});
</script>

<template>
  <div class="flex flex-col gap-5" v-if="sortedSubprocesses">
    <Subprocess v-for="s in sortedSubprocesses" :key="s.pid" :subprocess="s" />
  </div>
</template>

<style scoped></style>
