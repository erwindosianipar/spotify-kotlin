package com.example.kotlify.adapter

import android.content.Context
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ArrayAdapter
import android.widget.ImageView
import android.widget.TextView
import androidx.annotation.LayoutRes
import androidx.annotation.NonNull
import com.example.kotlify.R
import com.example.kotlify.model.Artist
import com.squareup.picasso.Picasso

class ArtistArrayAdapter(@NonNull context: Context, @LayoutRes layoutRes: Int=0, var artistList: MutableList<Artist>): ArrayAdapter<Artist>(context, layoutRes, artistList) {
    override fun getView(position: Int, convertView: View?, parent: ViewGroup): View {
        var itemView = LayoutInflater.from(context).inflate(R.layout.artist_item, parent, false)

        val artist = artistList.get(position)

        itemView?.findViewById<TextView>(R.id.id)?.setText(artist.id!!.toString())
        itemView?.findViewById<TextView>(R.id.name)?.setText(artist.name)
        itemView?.findViewById<TextView>(R.id.debut)?.setText(artist.debut)
        itemView?.findViewById<TextView>(R.id.category)?.setText(artist.category)

        val imageView= itemView?.findViewById<ImageView>(R.id.image)
        Picasso.get().load(artist.image).placeholder(R.drawable.ic_person_black_24dp).into(imageView)

        return itemView
    }
}