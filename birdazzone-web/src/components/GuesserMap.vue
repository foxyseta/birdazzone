<script setup lang="ts">
import ApiRepository from '../api/api-repository';
import type { Coordinates } from '../api/interfaces/tweet';
import { onBeforeMount, ref } from 'vue'
import ErrorWidget from './ErrorWidget.vue'
import FilterMap from './FilterList.vue'

const props = defineProps<{ gameId: string, key: number, from: string | null, to: string | null }>()

const ROME = [ 12.706374170037495, 42.21140846575139 ]

const error = ref<boolean>(false)
const errorTitle = ref<string>()
const errorText = ref<string>()

const center = ref<number[]>(ROME)
const zoom = ref<number>(6)
const rotation = ref(0)
const fillColor =ref( '#1eb980')
const strokeColor =ref( 'white')
const strokeWidth =ref(3)
const radius =ref(10)
const projection = ref("EPSG:4326")
const coordinates = ref<number[][]>([]) // [lat, long]
const from = ref<string | null>(props.from)
const to = ref<string | null>(props.to)

const unpackCoordinates = (coordinates: Coordinates) => [coordinates.latitude, coordinates.longitude]

const fetchCoordinates = async () => {
  if(!from.value || !to.value){
    const response = await ApiRepository.getListOfGuesser(props.gameId, "1", "100")
    if (response.esit) {
      // @ts-ignore
      coordinates.value = response.data!.entries.map(tweet => tweet.coordinates).filter(c => c).map(unpackCoordinates) // Remove undefined and nulls
    } else {
      if (response.statusCode === 204) {
      error.value = true
        errorTitle.value = 'There are no tweets!'
        errorText.value = 'No game instance has been played!'
      }
      else {
        error.value = true
        errorTitle.value = 'Error!'
        errorText.value = 'something went wrong!'
      }
    }
  }
  else{
    const response = await ApiRepository.getListOfGuesserFiltered(props.gameId, from.value.toString(), to.value.toString(), "1", "100")
    if (response.esit) {
      // @ts-ignore
      coordinates.value = response.data!.entries.map(tweet => tweet.coordinates).filter(c => c).map(unpackCoordinates) // Remove undefined and nulls
    } else {
      if (response.statusCode === 204) {
        error.value = true
        errorTitle.value = 'There are no tweets!'
        errorText.value = 'No game instance has been played!'
      }
      else {
        error.value = true
        errorTitle.value = 'Error!'
        errorText.value = 'something went wrong!'
      }
    }
  }
}

onBeforeMount(async () => {
  fetchCoordinates()
})

</script>

<template>
  <!-- Error -->
  <div v-if="error" class="flex justify-center items-center w-full">
    <ErrorWidget :open="true" :title="errorTitle" :text="errorText" />
  </div>
  
  <!-- Non error -->

<div class="flex">
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
