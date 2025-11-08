<script setup lang="ts">
import { Subprocess } from "../../model/subprocess";
import CopiesDetectors from "../copies_detector/CopiesDetectors.vue";
import SubprocessHeader from "./SubprocessHeader.vue";

const props = defineProps<{
  subprocess: Subprocess;
}>();
</script>

<template>
  <div class="shadow-lg border-2 border-gray-100 rounded-xl">
    <div class="border-b-2 border-b-gray-100 px-5 py-2">
      <SubprocessHeader :subprocess="subprocess" />
    </div>
    <div class="px-5 py-2">
      <CopiesDetectors :http-port="subprocess.httpPort" />
    </div>
  </div>
</template>

<style scoped></style>
