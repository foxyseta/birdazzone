<script setup lang=ts>
import { ref, onMounted } from 'vue'
import { createPopper } from "@popperjs/core";
const canvas = ref(null);

const success = 0.55;
const fail = 1-success;

onMounted(() => {
  const context = canvas.value.getContext('2d');
  const centerX = canvas.value.width / 2;
  const centerY = canvas.value.height / 2;
  const radius = 50;

  const spaceBetween = 0.33;
  const start = 3/2*Math.PI;

  const succPerc = (2 * Math.PI - spaceBetween) *success;     // 2PI : 1 =  x : success
  const failPerc = (2 * Math.PI - spaceBetween) - succPerc;

  context.beginPath();    // success
  context.lineCap = 'round';
  context.lineWidth = 13;
  context.arc(centerX, centerY, radius, start, start + succPerc, false);
  context.strokeStyle = '#1eb980';
  context.stroke();
  context.closePath();

  context.beginPath();    // fail
  context.arc(centerX, centerY, radius, start + succPerc + spaceBetween, start + succPerc + failPerc, false);
  context.strokeStyle = '#ff937f';
  context.lineWidth = 13;
  context.stroke();
  context.closePath();
});

const see = ref(false);
const popoverRef = ref(null);


const popover = () => {
  see.value = true;
  createPopper(canvas, popoverRef, {
    placement: 'top'
  });
  setTimeout(() => see.value = false, 2000);
}

</script>

<template>
  <div class="flex bg-foreground font-semibold text-lg rounded-lg h-full p-4 items-center place-self-center">
    <div class="items-center mx-3">
      <div class="text-white m-2">280 tried</div>
      <div class="m-2"><div class="text-lgreen inline">70 </div><div class="text-white inline">succeded</div></div>
      <div class="m-2"><div class="text-lred inline">210 </div><div class="text-white inline">failed</div></div>
    </div>
    <div class="relative">
      <canvas @mouseover="popover()" class="mr-3" ref="canvas" width="150" height="150"></canvas>
      <div ref="popoverRef" v-bind:class="{'hidden': !see, 'block': see}"
        class="absolute inset-0 text-sm font-semibold text-center">
        <div class="items-center pt-14 pr-3">
            <div class="text-lred">failed</div>
            <div class="text-lgreen">succeded</div>
        </div>
      </div>

      <div class="bottom-0 left-0 w-full ml-6">
        <div class="space-x-6">
          <div class="text-lred inline">{{ (fail * 100).toFixed(0) }}%</div>
          <div class="text-lgreen inline">{{ (success * 100).toFixed(0) }}%</div>
        </div>
      </div>
    </div>
  </div>
</template>
