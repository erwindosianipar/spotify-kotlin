package com.example.architecturecomponent.fragment

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.TextView
import androidx.fragment.app.Fragment
import androidx.lifecycle.Observer
import androidx.lifecycle.ViewModelProviders
import com.example.architecturecomponent.ArtistViewModel
import com.example.architecturecomponent.R
import com.example.architecturecomponent.model.Artist

class FragmentTwo: Fragment() {

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {

        val view: View = inflater.inflate(R.layout.fragment_two_layout, container, false)
        val artistViewModel = activity?.run {
            ViewModelProviders.of(this).get(ArtistViewModel::class.java)
        }

        artistViewModel?.artist?.observe(this, Observer {
            artist -> setArtist(artist)
        })
        return view
    }

    private fun setArtist(artist: Artist) {
        view?.findViewById<TextView>(R.id.artistNameTextView)?.setText(artist.name)
    }
}