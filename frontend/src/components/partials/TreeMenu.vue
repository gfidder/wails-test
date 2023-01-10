<script lang="ts" setup>
import { computed, ref } from "vue";
import PlusBoxOutline from "~icons/mdi/plus-box-outline";
import MinusBoxOutline from "~icons/mdi/minus-box-outline";
import { OidTree } from "../../utils/treeBuilder";

const showChildren = ref(false);

const props = defineProps<{
  node: OidTree;
  depth: number;
}>();

const indent = computed(() => {
  return { transform: `translate(${props.depth * 20}px)` };
});

function toggleChildren() {
  showChildren.value = !showChildren.value;
}

const cursorClass = computed(() => {
  if (props.node.children !== undefined) {
    return "cursor-pointer";
  } else {
    return "cursor-default";
  }
});

const hasChildren = computed(() => {
  return props.node.children !== undefined;
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
          {{ node.name }}
        </p>
      </div>
    </div>
    <div v-if="showChildren">
      <TreeMenu
        v-for="oid in node.children"
        :key="oid.oid"
        :node="oid"
        :depth="depth + 1"
      ></TreeMenu>
    </div>
  </div>
</template>
