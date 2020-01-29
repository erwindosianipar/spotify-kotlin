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
import com.example.kotlify.model.Genre
import com.example.kotlify.adapter.GenreArrayAdapter
import org.json.JSONArray
import org.json.JSONException

class GenreActivity: AppCompatActivity() {
    lateinit var listView: ListView
    lateinit var arrayAdapter: GenreArrayAdapter
    var url = "http://10.10.17.153:8080/genre"
    lateinit var requestQueue: RequestQueue
    var listGenre = mutableListOf<Genre>()

    var id = 0
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_genre)
        listView = findViewById(R.id.genre_list)

        id = intent.getIntExtra("ID", 0)
        arrayAdapter = GenreArrayAdapter(context = this, genreList = listGenre)
        requestQueue = Volley.newRequestQueue(this)
        listView.adapter = arrayAdapter

        listView.setOnItemClickListener { parent, view, position, id ->
            val newIntent = Intent(this, GenreActivityDetail::class.java).apply {
                putExtra("id", listGenre[position].id)
                putExtra("name", listGenre[position].name)
            }
            startActivity(newIntent)
        }

        fetchAllGenre()
    }

    private fun fetchAllGenre() {
        val genreRequest = StringRequest(Request.Method.GET, url,
            Response.Listener { response ->
                resolveSuccess(response)
            },
            Response.ErrorListener { error: VolleyError? ->
                Log.e("FETCH FAIL: ", error?.message)
            })
        requestQueue.add(genreRequest)
    }

    private fun resolveSuccess(response: String?) {
        try {
            val arrayResponse = JSONArray(response)
            for (i in 0 until arrayResponse.length()) {
                val item = arrayResponse.getJSONObject(i)
                val genre = Genre(
                    item.getInt("ID"),
                    item.getString("Name")
                )
                Log.i("Genre: $i", genre.toString())
                arrayAdapter.addAll(genre)
            }
        } catch (jsonEX: JSONException) {
            Log.e("PARSE", jsonEX.message)
        }
    }
}