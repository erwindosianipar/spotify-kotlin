package com.example.kotlify

import android.content.Intent
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.view.View

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
    }

    fun switchToSong(view: View) {
        var newIntent = Intent(this, SongActivity::class.java)
        startActivity(newIntent)
    }

    fun switchToGenre(view: View) {
        var newIntent = Intent(this, GenreActivity::class.java)
        startActivity(newIntent)
    }

    fun switchToArtist(view: View) {
        var newIntent = Intent(this, ArtistActivity::class.java)
        startActivity(newIntent)
    }
}
