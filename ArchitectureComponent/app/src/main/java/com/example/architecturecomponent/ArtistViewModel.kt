package com.example.architecturecomponent

import android.util.Log
import androidx.lifecycle.MutableLiveData
import androidx.lifecycle.ViewModel
import com.example.architecturecomponent.api.ArtistAPI
import com.example.architecturecomponent.config.RetrofitBuilder
import com.example.architecturecomponent.model.Artist
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response

class ArtistViewModel: ViewModel() {

    val artist: MutableLiveData<Artist> by lazy {
        MutableLiveData<Artist>()
    }

    fun getArtist(id: String) {
        val retrofit = RetrofitBuilder.createRetrofit()

        val artistAPI = retrofit.create(ArtistAPI::class.java)

        artistAPI.getArtist(id).enqueue(object : Callback<Artist>{
            override fun onFailure(call: Call<Artist>, t: Throwable) {
                Log.e("FETCH FAIL", t.message)
            }

            override fun onResponse(call: Call<Artist>, response: Response<Artist>) {
                Log.i("FETCH DONE", response.body().toString())
                artist.value=response.body()
            }
        })
    }
}