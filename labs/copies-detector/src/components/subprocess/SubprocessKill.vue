<script setup lang="ts">
import { useKill } from '../../mutations/subprocess';

const props = defineProps<{
    pid: number,
}>()

const { mutate } = useKill()

const onClick = () => {
    mutate(props.pid)
}
</script>

<template>
  <div>
    <button
      type="button"
      class="rounded-xl p-2 grid place-items-center text-2xl text-white bg-red-500 hover:cursor-pointer"
      @click="onClick"
    >
      <i class="pi pi-ban"></i>
    </button>
  </div>
</template>

<style scoped></style>
