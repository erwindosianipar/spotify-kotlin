package com.example.kotlify

import android.content.Intent
import android.os.Bundle
import android.util.Log
import android.widget.ListView
import androidx.appcompat.app.AppCompatActivity
import com.android.volley.Request
import com.android.volley.RequestQueue
import com.android.volley.Response
import com.android.volley.VolleyError
import com.android.volley.toolbox.StringRequest
import com.android.volley.toolbox.Volley
import com.example.kotlify.adapter.ArtistArrayAdapter
import com.example.kotlify.model.Artist
import org.json.JSONArray
import org.json.JSONException

class ArtistActivity: AppCompatActivity() {
    lateinit var listView: ListView
    lateinit var arrayAdapter: ArtistArrayAdapter
    var url = "http://10.10.17.153:8080/artist"
    lateinit var requestQueue: RequestQueue
    var listArtist = mutableListOf<Artist>()

    var id = 0
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_artist)

        listView = findViewById(R.id.artist_list)
        id = intent.getIntExtra("ID", 0)
        arrayAdapter = ArtistArrayAdapter(context = this, artistList = listArtist)
        requestQueue = Volley.newRequestQueue(this)
        listView.adapter = arrayAdapter

        listView.setOnItemClickListener { parent, view, position, id ->
            val newIntent = Intent(this, ArtistActivityDetail::class.java).apply {
                putExtra("id", listArtist[position].id)
                putExtra("name", listArtist[position].name)
            }
            startActivity(newIntent)
        }

        fetchAllArtist()
    }

    private fun fetchAllArtist() {
        val artistRequest = StringRequest(
            Request.Method.GET, url,
            Response.Listener { response ->
                resolveSuccess(response)
            },
            Response.ErrorListener { error: VolleyError? ->
                Log.e("FETCH FAIL: ", error?.message)
            })
        requestQueue.add(artistRequest)
    }

    private fun resolveSuccess(response: String?) {
        try {
            val arrayResponse = JSONArray(response)
            for (i in 0 until arrayResponse.length()) {
                val item = arrayResponse.getJSONObject(i)
                val artist = Artist(
                    item.getInt("ID"),
                    item.getString("Name"),
                    item.getString("Debut"),
                    item.getString("Category"),
                    item.getString("Image")
                )
                Log.i("Artist: $i", artist.toString())
                arrayAdapter.addAll(artist)
            }
        } catch (jsonEX: JSONException) {
            Log.e("PARSE", jsonEX.message)
        }
    }
}