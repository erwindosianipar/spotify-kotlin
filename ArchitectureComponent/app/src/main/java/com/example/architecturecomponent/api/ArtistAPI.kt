package com.example.architecturecomponent.api

import com.example.architecturecomponent.model.Artist
import retrofit2.Call
import retrofit2.http.GET
import retrofit2.http.Path

interface ArtistAPI {

    @GET("artist/{id}")
    fun getArtist(@Path("id") id: String): Call<Artist>

}