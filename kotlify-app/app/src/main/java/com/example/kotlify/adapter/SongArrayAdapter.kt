package com.example.kotlify.adapter

import android.content.Context
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ArrayAdapter
import android.widget.TextView
import androidx.annotation.LayoutRes
import androidx.annotation.NonNull
import com.example.kotlify.model.Song
import com.example.kotlify.R

class SongArrayAdapter(@NonNull context: Context, @LayoutRes layoutRes: Int=0, var songList: MutableList<Song>): ArrayAdapter<Song>(context, layoutRes, songList) {
    override fun getView(position: Int, convertView: View?, parent: ViewGroup): View {
        var itemView = LayoutInflater.from(context).inflate(R.layout.song_item, parent, false)

        val song = songList.get(position)

        itemView?.findViewById<TextView>(R.id.id)?.setText(song.id!!.toString())
        itemView?.findViewById<TextView>(R.id.title)?.setText(song.title)

        return itemView
    }
}