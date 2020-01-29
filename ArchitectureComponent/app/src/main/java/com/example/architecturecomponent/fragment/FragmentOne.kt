package com.example.architecturecomponent.fragment

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import android.widget.EditText
import androidx.fragment.app.Fragment
import androidx.lifecycle.ViewModelProviders
import com.example.architecturecomponent.ArtistViewModel
import com.example.architecturecomponent.R

class FragmentOne: Fragment() {

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {

        val view: View = inflater.inflate(R.layout.fragment_one_layout, container, false)
        val artistViewModel = activity?.run {
            ViewModelProviders.of(this).get(ArtistViewModel::class.java)
        }

        view.findViewById<Button>(R.id.frg1_fetch_button)
            .setOnClickListener{
                val id = view.findViewById<EditText>(R.id.frg1_id_text).text.toString()
                artistViewModel?.getArtist(id)
        }

        return view
    }
}