package com.example.kotlify

import android.os.Bundle
import android.util.Log
import android.widget.ListView
import android.widget.TextView
import androidx.appcompat.app.AppCompatActivity
import com.android.volley.Request
import com.android.volley.RequestQueue
import com.android.volley.Response
import com.android.volley.VolleyError
import com.android.volley.toolbox.StringRequest
import com.android.volley.toolbox.Volley
import com.example.kotlify.adapter.SongArrayAdapter
import com.example.kotlify.model.Song
import org.json.JSONArray
import org.json.JSONException

class GenreActivityDetail: AppCompatActivity() {
    lateinit var listView: ListView
    lateinit var arrayAdapter: SongArrayAdapter
    var url = "http://10.10.17.153:8080/song/genre/"
    lateinit var requestQueue: RequestQueue
    var listSong = mutableListOf<Song>()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_genre_detail)

        var id = intent.getIntExtra("id", 0)
        var name = intent.getStringExtra("name")

        name = "Genre $name"
        val viewName = findViewById<TextView>(R.id.genre_name)
        viewName.setText(name)

        listView = findViewById(R.id.song_list)
        arrayAdapter = SongArrayAdapter(context = this, songList = listSong)
        requestQueue = Volley.newRequestQueue(this)
        listView.adapter = arrayAdapter

        fetchSongByGenreID(id)
    }

    private fun fetchSongByGenreID(id: Int) {
        val songRequest = StringRequest(
            Request.Method.GET, url+id,
            Response.Listener { response ->
                resolveSuccess(response)
            },
            Response.ErrorListener { error: VolleyError? ->
                Log.e("FETCH FAIL: ", error?.message)
            })
        requestQueue.add(songRequest)
    }

    private fun resolveSuccess(response: String?) {
        try {
            val arrayResponse = JSONArray(response)
            for (i in 0 until arrayResponse.length()) {
                val item = arrayResponse.getJSONObject(i)
                val song = Song(
                    item.getInt("ID"),
                    item.getString("Name")
                )
                Log.i("Genre: $i", song.toString())
                arrayAdapter.addAll(song)
            }
        } catch (jsonEX: JSONException) {
            Log.e("PARSE", jsonEX.message)
        }
    }
}