<script lang="ts" setup>
import { onBeforeMount, ref } from "vue";
import ApiRepository from "@/api/api-repository";
import type { Coordinates } from "@/api/interfaces/tweet";

const ROME = [ 18.706374170037495, 42.21140846575139 ]

const props = defineProps<{gameId: string}>()
const coordinates = ref<Coordinates[]>([])
const center = ref<number[]>(ROME)
const zoom = ref<number>(6)
const rotation = ref(0)
const coordinate =ref(ROME)
const fillColor =ref( 'red')
const strokeColor =ref( 'white')
const strokeWidth =ref( 10)
const radius =ref( 40)


const unpackCoordinates = (coordinates: Coordinates) => [coordinates.latitude, coordinates.longitude]

const fetchCoordinates = async () => {
  const response = await ApiRepository.getListOfGuesser(props.gameId)
  if (response.esit) {
    coordinates.value = response.data!.entries.map(tweet => tweet.coordinates).filter(c => c) // Remove undefined and nulls
    console.log(coordinates.value)
  } else {

  }
}

onBeforeMount(async () => {
  //fetchCoordinates()
})

</script>
<template>
  <div class="bg-foreground rounded-xl w-100 p-3">
    <ol-map :loadTilesWhileAnimating="true" :loadTilesWhileInteracting="true" style="height:400px">

    <ol-view :center="center" :rotation="rotation" :zoom="zoom" />

    <OlTileLayer>
        <OlSourceOsm/>
    </OlTileLayer>

            <ol-vector-layer>
        <ol-source-vector>
            <ol-feature>
                <ol-geom-point :coordinates="coordinate"></ol-geom-point>
                <ol-style>
                    <ol-style-circle :radius="radius">
                        <ol-style-fill :color="fillColor"></ol-style-fill>
                        <ol-style-stroke :color="strokeColor" :width="strokeWidth"></ol-style-stroke>
                    </ol-style-circle>
                </ol-style>
            </ol-feature>

        </ol-source-vector>

    </ol-vector-layer>

</ol-map>
  </div>
</template>
