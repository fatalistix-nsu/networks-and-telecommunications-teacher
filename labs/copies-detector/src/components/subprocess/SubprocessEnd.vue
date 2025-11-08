<script setup lang="ts">
import { useEnd } from '../../mutations/subprocess';

const props = defineProps<{
	pid: number
}>()

const { mutate } = useEnd()

const onClick = () => {
  console.log("ending!!!: " + props.pid)
	mutate(props.pid)	
};
</script>

<template>
  <div>
    <button
      type="button"
      class="bg-orange-400 rounded-xl p-2 text-2xl grid place-items-center text-white hover:cursor-pointer"
      @click="onClick"
    >
      <i class="pi pi-stop"></i>
    </button>
  </div>
</template>

<style scoped></style>
