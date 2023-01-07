<script lang="ts" setup>
import { computed, ref } from "vue";
import PlusBoxOutline from "~icons/mdi/plus-box-outline";
import MinusBoxOutline from "~icons/mdi/minus-box-outline";

const showChildren = ref(false);

const props = defineProps<{
  label: string;
  nodes?: any[];
  depth: number;
}>();

const indent = computed(() => {
  return { transform: `translate(${props.depth * 20}px)` };
});

function toggleChildren() {
  showChildren.value = !showChildren.value;
}

const cursorClass = computed(() => {
  if (props.nodes !== undefined) {
    return "cursor-pointer";
  } else {
    return "cursor-default";
  }
});

const hasChildren = computed(() => {
  return props.nodes !== undefined;
});
</script>

<template>
  <div>
    <div class="pb-1 mb-1" @click="toggleChildren">
      <div :style="indent" :class="cursorClass" class="flex text-gray-50">
        <PlusBoxOutline v-if="hasChildren && !showChildren" />
        <MinusBoxOutline v-else-if="hasChildren && showChildren" />
        <!--TODO : add padding here so everything lines up even if there is no child type icon -->
        <p class="pl-1">
          {{ label }}
        </p>
      </div>
    </div>
    <TreeMenu
      v-if="showChildren"
      v-for="node in nodes"
      :nodes="node.nodes"
      :label="node.label"
      :depth="depth + 1"
    ></TreeMenu>
  </div>
</template>
