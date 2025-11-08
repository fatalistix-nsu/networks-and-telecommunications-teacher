<script setup lang="ts">
import { useRun } from '../../mutations/subprocess';

const { mutate } = useRun()

const onClick = () => {
  mutate()
}
</script>

<template>
  <div>
    <button
      type="button"
      class="flex justify-center gap-1 items-center border-2 border-dashed rounded-xl p-3 w-full transition-all duration-300 text-gray-400 hover:bg-linear-to-r hover:from-purple-400 hover:to-orange-400 hover:text-white hover:outline-2 hover:outline-white hover:outline-offset-[-2px] hover:cursor-pointer"
      @click="onClick"
    >
        <i class="pi pi-plus"></i>
        Spawn subprocess
    </button>
  </div>
</template>

<style scoped></style>
