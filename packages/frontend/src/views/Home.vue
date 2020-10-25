<template>
  <v-container>
    <v-row>
      <v-col class="col-6">
        <v-card elevation="2">
          <v-card-title>Compteur de visite</v-card-title>
          <v-card-text>
            Nombre de visite actuelle : {{ visitCount }}
          </v-card-text>
        </v-card>
      </v-col>
      <v-col class="col-6">
        <v-card elevation="2">
          <v-card-title>Nom enregistrés</v-card-title>
          <v-simple-table>
            <template #default>
              <thead>
                <tr>
                  <th class="text-left">
                    Mongo ObjectID
                  </th>
                  <th class="text-left">
                    Name
                  </th>
                </tr>
              </thead>
              <tbody>
                <template v-if="names.length > 0">
                  <tr v-for="item in names" :key="item.name">
                    <td>{{ item.ID }}</td>
                    <td>{{ item.name }}</td>
                  </tr>
                </template>
                <template v-else>
                  <tr>
                    <td colspan="2">Aucuns noms enregistrés</td>
                  </tr>
                </template>
              </tbody>
            </template>
          </v-simple-table>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Prop, Vue } from "vue-property-decorator";
import axios from "axios";

@Component
export default class Home extends Vue {
  @Prop({ required: false })
  private putName!: string | null;
  private visitCount = 0;
  private names: Array<{ ID: string; name: string }> = [];

  async mounted(): Promise<void> {
    if (null !== this.putName) {
      await this.doPutName();
    }

    this.fetchVisitCount();
    this.fetchNames();
  }

  private async fetchVisitCount(): Promise<void> {
    try {
      const response = await axios.get("/api/visit");
      this.visitCount = response.data.count;
    } catch (e) {
      console.error(e.message);
    }
  }

  private async fetchNames(): Promise<void> {
    try {
      const response = await axios.get("/api/people");
      this.names = response.data.people;
    } catch (e) {
      console.error(e.message);
    }
  }

  private async doPutName(): Promise<void> {
    try {
      await axios.post("/api/people", { name: this.putName });
    } catch (e) {
      console.error(e.message);
    }
  }
}
</script>
