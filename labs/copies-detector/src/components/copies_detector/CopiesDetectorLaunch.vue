<script setup lang="ts">
import { useCreateCopiesDetector } from "../../mutations/copiesDetector";

const props = defineProps<{
  httpPort: number;
  host: string;
  port: number;
  name: string;
}>();

const { mutate } = useCreateCopiesDetector(props.httpPort);

const onClick = () => {
  mutate({
    host: props.host,
    port: props.port,
    name: props.name,
  });
};
</script>

<template>
  <button
    type="button"
    class="w-full h-full flex items-center justify-center rounded-xl p-2 text-white bg-green-500 hover:cursor-pointer"
    @click="onClick"
  >
    <i class="pi pi-caret-right"></i>
    Launch
  </button>
</template>

<style scoped></style>
