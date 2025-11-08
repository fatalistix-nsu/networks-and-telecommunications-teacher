<script setup lang="ts">
const copiesText = [
  "  /$$$$$$                      /$$                    ",
  " /$$__  $$                    |__/                    ",
  "| $$  \\__/  /$$$$$$   /$$$$$$  /$$  /$$$$$$   /$$$$$$$",
  "| $$       /$$__  $$ /$$__  $$| $$ /$$__  $$ /$$_____/",
  "| $$      | $$  \\ $$| $$  \\ $$| $$| $$$$$$$$|  $$$$$$ ",
  "| $$    $$| $$  | $$| $$  | $$| $$| $$_____/ \\____  $$",
  "|  $$$$$$/|  $$$$$$/| $$$$$$$/| $$|  $$$$$$$ /$$$$$$$/",
  " \\______/  \\______/ | $$____/ |__/ \\_______/|_______/ ",
  "                    | $$                              ",
  "                    | $$                              ",
  "                    |__/                              ",
];

const detectorText = [
  "       /$$             /$$                           /$$                        ",
  "      | $$            | $$                          | $$                        ",
  "  /$$$$$$$  /$$$$$$  /$$$$$$    /$$$$$$   /$$$$$$$ /$$$$$$    /$$$$$$   /$$$$$$ ",
  " /$$__  $$ /$$__  $$|_  $$_/   /$$__  $$ /$$_____/|_  $$_/   /$$__  $$ /$$__  $$",
  "| $$  | $$| $$$$$$$$  | $$    | $$$$$$$$| $$        | $$    | $$  \\ $$| $$  \\__/",
  "| $$  | $$| $$_____/  | $$ /$$| $$_____/| $$        | $$ /$$| $$  | $$| $$      ",
  "|  $$$$$$$|  $$$$$$$  |  $$$$/|  $$$$$$$|  $$$$$$$  |  $$$$/|  $$$$$$/| $$      ",
  " \\_______/ \\_______/   \\___/   \\_______/ \\_______/   \\___/   \\______/ |__/      ",
];
</script>

<template>
  <div class="flex flex-col mt-5">
    <div class="header-copies-grid">
      <div class="header-copies-grid-element">
        <h1
          class="font-mono whitespace-pre p-0 m-0"
          v-for="line in copiesText"
        >
          {{ line }}
        </h1>
      </div>
    </div>
    <div class="header-detector-grid">
      <div class="header-detector-grid-element">
        <h1
          class="font-mono whitespace-pre p-0 m-0"
          v-for="line in detectorText"
        >
          {{ line }}
        </h1>
      </div>
    </div>
  </div>
</template>

<style scoped>
@reference "../index.css";

.header-copies-grid {
  @apply grid grid-cols-1;

  @variant lg {
    @apply grid-cols-10;
  }

  @variant xl {
    @apply grid-cols-12;
  }
}

.header-copies-grid-element {
  @apply col-start-1 row-start-1;

  @variant lg {
    @apply col-start-2;
  }

  @variant xl {
    @apply col-start-3;
  }

  @variant 2xl {
    @apply col-start-3;
  }
}

.header-detector-grid {
  @apply grid grid-cols-1;

  @variant lg {
    @apply grid-cols-10 grid-flow-col-dense justify-items-end;
  }

  @variant xl {
    @apply grid-cols-12 grid-flow-col-dense justify-items-end;
  }
}

.header-detector-grid-element {
  @apply col-start-1 row-start-1;

  @variant lg {
    @apply col-start-9;
  }

  @variant xl {
    @apply col-start-10;
  }

  @variant 2xl {
    @apply col-start-10;
  }
}
</style>
