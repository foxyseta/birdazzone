<script setup lang="ts">
import ApiRepository from '@/api/api-repository';
import type { Coordinates } from '@/api/interfaces/tweet';
import { onBeforeMount, ref } from 'vue'
import FilterMap from './FilterList.vue'

const props = defineProps<{gameId: string}>()

const ROME = [ 12.706374170037495, 42.21140846575139 ]

const center = ref<number[]>(ROME)
const zoom = ref<number>(6)
const rotation = ref(0)
const fillColor =ref( '#1eb980')
const strokeColor =ref( 'white')
const strokeWidth =ref(3)
const radius =ref(10)
const projection = ref("EPSG:4326")
const coordinates = ref<number[][]>([]) // [lat, long]
const from = ref<string | null>()
const to = ref<string | null>()

const unpackCoordinates = (coordinates: Coordinates) => [coordinates.latitude, coordinates.longitude]

const fetchCoordinates = async () => {
  if(!from.value || !to.value){
    const response = await ApiRepository.getListOfGuesser(props.gameId, "0", "100")
    if (response.esit) {
      // @ts-ignore
      coordinates.value = response.data!.entries.map(tweet => tweet.coordinates).filter(c => c).map(unpackCoordinates) // Remove undefined and nulls
      console.log(response.data!.entries.map(tweeet => tweeet.coordinates))
      coordinates.value.push(ROME)
      coordinates.value.push([ROME[0], ROME[1] +2])
      coordinates.value.push([ROME[0] +0.1, ROME[1] +0.1])
      console.log(coordinates.value)
    } else {

    }
  }
  else{
    const response = await ApiRepository.getListOfGuesserFiltered(props.gameId, from.value.toString(), to.value.toString(), "0", "100")
    if (response.esit) {
      // @ts-ignore
      coordinates.value = response.data!.entries.map(tweet => tweet.coordinates).filter(c => c).map(unpackCoordinates) // Remove undefined and nulls
      console.log(response.data!.entries.map(tweeet => tweeet.coordinates))
      coordinates.value.push(ROME)
      coordinates.value.push([ROME[0], ROME[1] + 2])
      coordinates.value.push([ROME[0] + 0.1, ROME[1] + 0.1])
      console.log(coordinates.value)
    } else {

    }
  }
}

onBeforeMount(async () => {
  fetchCoordinates()
})

</script>

<template>

<div class="flex">
  <div class="flex flex-col m-6 mr-8 ml-20">
    <FilterMap :from="from" :to="to" @change-from-to="(f, t) => {from = f; to = t; fetchCoordinates()}" />
  </div>

  <div class="flex mt-6 p-5 bg-foreground shadow rounded-xl justify-center">
    <ol-map :loadTilesWhileAnimating="true" :loadTilesWhileInteracting="true" style="height: 50rem; width: 50rem">

      <ol-zoom-control />

    <ol-view ref="view" :center="center" :rotation="rotation" :zoom="zoom" :projection="projection"/>

    <ol-tile-layer>
        <ol-source-osm />
    </ol-tile-layer>

          <ol-vector-layer>
            <ol-source-vector>
                <ol-feature>
                    <ol-geom-multi-point :coordinates="coordinates"></ol-geom-multi-point>
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
</div>
</template>
