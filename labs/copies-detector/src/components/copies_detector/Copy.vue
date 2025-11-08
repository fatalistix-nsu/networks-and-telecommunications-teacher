<script setup lang="ts">
import { onBeforeUnmount, onMounted, ref } from "vue";
import { useGetCopiesDetectorService } from "../../services/copiesDetector";
import { DetectedCopy } from "../../model/copiesDetector";
import { sleep } from "../../utils/time";

const props = defineProps<{
  id: string;
  httpPort: number;
}>();

const service = useGetCopiesDetectorService(props.httpPort);
const detectedCopies = ref<DetectedCopy[]>([]);
const now = ref<Date>(new Date());
const running = ref(false);

const refreshLoop = async () => {
  while (running) {
    now.value = new Date();

    service.run(props.id).then((v) => {
      detectedCopies.value = v.detectedCopies;
    });

    await sleep(1_000);
  }
};

onMounted(() => {
  running.value = true;
  refreshLoop();
});

onBeforeUnmount(() => {
  running.value = false;
});
</script>

<template>
  <div v-for="v in detectedCopies" :key="v.name">
    {{
      v.name +
      " " +
      v.lastRefresh.getHours().toString().padStart(2, "0") +
      ":" +
      v.lastRefresh.getMinutes().toString().padStart(2, "0") +
      ":" +
      v.lastRefresh.getSeconds().toString().padStart(2, "0") +
      "." +
      v.lastRefresh.getMilliseconds().toString().padStart(3, "0")
    }}
  </div>
</template>

<style scoped></style>
