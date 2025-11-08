<script setup lang="ts">
import CopyList from "./CopyList.vue";
import CopiesDetectorHeader from "./CopiesDetectorHeader.vue";

const props = defineProps<{
  httpPort: number,
}>()

</script>

<template>
  <div class="py-2">
    <div>
      <CopiesDetectorHeader :http-port="httpPort" />
    </div>
    <div>
      <CopyList :http-port="httpPort" />
    </div>
  </div>
</template>

<style scoped></style>
