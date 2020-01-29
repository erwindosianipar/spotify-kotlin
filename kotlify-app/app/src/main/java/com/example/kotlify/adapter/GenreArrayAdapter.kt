package com.example.kotlify.adapter

import android.content.Context
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ArrayAdapter
import android.widget.TextView
import androidx.annotation.LayoutRes
import androidx.annotation.NonNull
import com.example.kotlify.model.Genre
import com.example.kotlify.R

class GenreArrayAdapter(@NonNull context: Context, @LayoutRes layoutRes: Int=0, var genreList: MutableList<Genre>): ArrayAdapter<Genre>(context, layoutRes, genreList) {
    override fun getView(position: Int, convertView: View?, parent: ViewGroup): View {
        var itemView = LayoutInflater.from(context).inflate(R.layout.genre_item, parent, false)

        val genre = genreList.get(position)

        itemView?.findViewById<TextView>(R.id.id)?.setText(genre.id!!.toString())
        itemView?.findViewById<TextView>(R.id.name)?.setText(genre.name)

        return itemView
    }
}