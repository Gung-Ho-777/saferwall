<template>
  <div class="tile is-child box">
    <h4 class="title">Basic Properties</h4>
    <div v-for="(i, index) in Properties" class="data-data" :key="index">
      <strong class="data-label">
        {{ getLabelForGivenKey(index) }}
      </strong>
      <!-- TRiD -->
      <span
        class="data-value"
        :class="{ 'trid-container': index === 'trid' }"
        v-if="index === 'trid'"
      >
        <p v-for="(t, index) in summaryData.trid" :key="index">
          <span class="trid">
            <span class="value-text">{{ t }}</span>

            <copy :content="t"></copy>
          </span>
        </p>
      </span>
      <!-- Packer -->
      <span
        class="data-value"
        :class="{ 'packer-container': index === 'packer' }"
        v-else-if="index === 'packer'"
      >
        <p v-for="(t, index) in summaryData.packer" :key="index">
          <span class="packer">
            <span class="value-text">{{ t }}</span>

            <copy :content="t"></copy>
          </span>
        </p>
      </span>
      <!-- Tags -->
      <span class="tags" v-else-if="index === 'tags'">
        <span
          v-for="(item, index) in summaryData.tags"
          :key="index"
          class="tag is-link is-normal"
        >
          <span>{{ item }}</span>
          <copy :content="item"></copy>
        </span>
      </span>
      <!-- Size -->
      <span class="data-value" v-else-if="index === 'size'">
          <span class="value-text">{{bytesToSize(i)}} ({{i}})</span>
          <copy :content="i"></copy>
      </span>
      <!-- Default -->
      <span class="data-value" v-else>
        <span class="value-text">{{
          index !== "sha512" ? i : i.substring(0, 70) + "..."
        }}</span>
        <copy :content="i"></copy>
      </span>
    </div>
  </div>
</template>

<script>
import Copy from "@/components/elements/Copy"

export default {
  props: ["summaryData"],
  components: {
    copy: Copy,
  },
  data() {
    return {
      uppercaseFields: ["md5", "sha-1", "sha-256", "sha-512", "crc32"],
    }
  },
  computed: {
    Properties: function() {
      return this._.pick(this.summaryData, [
        "size",
        "crc32",
        "md5",
        "sha1",
        "sha256",
        "sha512",
        "ssdeep",
        "magic",
        "packer",
        "trid",
        "tags",
      ])
    },
  },
  methods: {
    getLabelForGivenKey(key) {
      return this.uppercaseFields.includes(key)
        ? key.toUpperCase()
        : key === "filesize"
        ? "File Size"
        : key === "trid"
        ? "TRiD"
        : key === "ssdeep"
        ? "SSDeep"
        : key
    },
    bytesToSize(bytes) {
      var sizes = ["Bytes", "KB", "MB", "GB", "TB"]
      if (bytes === 0) return "0 Byte"
      var i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)))
      return Math.round(bytes / Math.pow(1024, i), 2) + " " + sizes[i]
    },
  },
}
</script>

<style></style>
